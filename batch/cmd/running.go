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
			}

			// SSH Key を払い出す
			keys, err := NewKeys()
			if err != nil {
				fmt.Println(err)
			}
			privateKeyFilePath, err := keys.CreatePrivateKeyFile(instanceID)
			if err != nil {
				fmt.Println(err)
			}
			publicKeyFilePath, data, err := keys.CreatePublicKeyFile(instanceID)
			if err != nil {
				fmt.Println(err)
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
			var out []byte
			out, err = exec.Command("docker", fmt.Sprintf(
				"run -d --cap-add=SYS_ADMIN -v /sys/fs/cgroup:/sys/fs/cgroup:ro local/c8-systemd-ssh --name instance-%v --hostname instance-%v",
				instanceID, instanceID)).Output()
			if err != nil {
				fmt.Println(err)
			}
			log.Println(out)

			// TODO: Key の出力先を変える
			out, err = exec.Command("docker", fmt.Sprintf(
				"cp instance-%v:/root/.ssh/id_rsa /tmp/instance-%v-private.pem",
				instanceID, instanceID)).CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}
			log.Println(out)

			out, err = exec.Command("docker", fmt.Sprintf(
				"cp instance-%v:/root/.ssh/id_rsa /tmp/instance-%v-public.pem",
				instanceID, instanceID)).CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}
			log.Println(out)

			// Networkをつくる
			out, err = exec.Command("docker", fmt.Sprintf(
				"network create --driver=bridge --subnet=%v.0/24 --gateway=%v1 -o “com.docker.network.bridge.name=br_nic1” mybridge-%v",
				ip[:strings.LastIndex(ip, ".")], ip, instanceID)).CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}
			log.Println(out)

			// Container と Networkを紐付ける
			out, err = exec.Command("docker", fmt.Sprintf(
				"network connect --ip=%v mybridge-%v instance-%v",
				ip, instanceID, instanceID)).CombinedOutput()
			if err != nil {
				fmt.Println(err)
			}
			log.Println(out)

			// Instance を Running にする
			if _, err := instanceStateUpdate(instanceID, StateRunning); err != nil {
				fmt.Println(err)
			}

			// TODO: Private Key と設定情報を渡す
			fmt.Printf("Instance#%v running\n", instanceID)
		}
	},
}

func init() {
	rootCmd.AddCommand(runningCmd)
}
