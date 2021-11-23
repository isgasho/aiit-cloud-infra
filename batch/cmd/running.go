package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// runningCmd represents the running command
var runningCmd = &cobra.Command{
	Use:   "running",
	Short: "run instance",
	Long:  `run instance`,
	Run: func(cmd *cobra.Command, args []string) { // TODO: cobra error handling
		res, err := getInstances(StateStarting)
		if err != nil {
			fmt.Println(err)
		}

		for _, row := range res {
			instanceID := row.ID
			ip := row.IPAddress

			// Instance を Initializing にする
			if _, err := instanceStateUpdate(instanceID, StateInitializing); err != nil {
				fmt.Println(err)
				return
			}

			// SSH Key を払い出す
			keys, err := NewKeys()
			if err != nil {
				fmt.Println(err)
				return
			}
			privateKeyFilePath, err := keys.CreatePrivateKeyFile(instanceID)
			if err != nil {
				fmt.Println(err)
				return
			}
			publicKeyFilePath, data, err := keys.CreatePublicKeyFile(instanceID)
			if err != nil {
				fmt.Println(err)
				return
			}

			log.Printf("private: %v\n", privateKeyFilePath)
			log.Printf("public: %v\n", publicKeyFilePath)
			log.Println(data) //keys.data

			// TODO: Database の keys.data を更新する

			// Instance を作る
			log.Printf("size: %v", row.Size)
			log.Printf("IP address: %v", ip)
			log.Printf("MAC address: %v", row.MacAddress)

			// Container を作る
			// TODO: Docker API を利用する
			if err := exec.Command("docker", fmt.Sprintf(
				"run -d --name instance-%v --hostname instance-%v --cap-add=SYS_ADMIN -v /sys/fs/cgroup:/sys/fs/cgroup:ro local/c8-systemd-ssh",
				instanceID, instanceID)).Run(); err != nil {
				log.Printf("docker run -d --name instance-%v --hostname instance-%v --cap-add=SYS_ADMIN -v /sys/fs/cgroup:/sys/fs/cgroup:ro local/c8-systemd-ssh\n",
					instanceID, instanceID)
				fmt.Println(err)
			}

			// TODO: Key の出力先を変える
			if err := exec.Command("docker", fmt.Sprintf(
				"cp instance-%v:/root/.ssh/id_rsa /tmp/instance-%v-private.pem",
				instanceID, instanceID)).Run(); err != nil {
				log.Printf("docker cp instance-%v:/root/.ssh/id_rsa /tmp/instance-%v-private.pem\n",
					instanceID, instanceID)
				fmt.Println(err)
			}

			if err := exec.Command("docker", fmt.Sprintf(
				"cp instance-%v:/root/.ssh/id_rsa /tmp/instance-%v-public.pem",
				instanceID, instanceID)).Run(); err != nil {
				log.Printf("docker cp instance-%v:/root/.ssh/id_rsa /tmp/instance-%v-public.pem\n",
					instanceID, instanceID)
				fmt.Println(err)
			}

			// Networkをつくる
			if err := exec.Command("docker", fmt.Sprintf(
				"network create --driver=bridge --subnet=%v.0/24 --gateway=%v -o “com.docker.network.bridge.name=br_nic1” mybridge-%v",
				ip[:strings.LastIndex(ip, ".")], ip, instanceID)).Run(); err != nil {
				log.Printf("docker network create --driver=bridge --subnet=%v.0/24 --gateway=%v -o “com.docker.network.bridge.name=br_nic1” mybridge-%v\n",
					ip[:strings.LastIndex(ip, ".")], ip, instanceID)
				fmt.Println(err)
			}

			// Container と Networkを紐付ける
			if err := exec.Command("docker", fmt.Sprintf(
				"network connect --ip=%v mybridge-%v instance-%v",
				ip, instanceID, instanceID)).Run(); err != nil {
				log.Printf("docker network connect --ip=%v mybridge-%v instance-%v\n",
					ip, instanceID, instanceID)
				fmt.Println(err)
			}

			// Instance を Running にする
			if _, err := instanceStateUpdate(instanceID, StateRunning); err != nil {
				fmt.Println(err)
				return
			}

			// TODO: Private Key と設定情報を渡す
			log.Printf("chmod 600 %v", publicKeyFilePath)
			log.Printf("ssh -i /tmp/instance-%v-private.pem root@%v", instanceID, ip)
			fmt.Printf("Instance#%v running\n", instanceID)
		}
	},
}

func init() {
	rootCmd.AddCommand(runningCmd)
}
