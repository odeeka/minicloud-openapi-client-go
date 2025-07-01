package main

import (
	"fmt"
	"log"

	"github.com/odeeka/minicloud-openapi-client-go/minicloud"
)

func main() {
	client, err := minicloud.NewClient("http://localhost:8080", "testuser", "TestPass2025")
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	fmt.Println("Client successfully created\nToken:", client.Token)

	// Example for listing (data source) resources
	vms, err := client.ListVMs()
	if err != nil {
		log.Fatalf("Failed to list VMs: %v", err.Error())
	}

	for _, vm := range vms {
		fmt.Printf("VM ID: %d, Name: %s, Image: %s, CPU: %f, Memory: %d, Ports: %v, Env: %v, ContainerID: %s\n", vm.ID, vm.Name, vm.Image, vm.CPU, vm.Memory, vm.Ports, vm.Env, vm.ContainerID)
	}

	storages, err := client.ListStorages()
	if err != nil {
		log.Fatalf("failed to list storages: %v", err)
	}

	for _, s := range storages {
		fmt.Printf("Storage ID: %d, Name: %s, Size: %d GB\n", s.ID, s.Name, s.SizeGB)
	}

	// // Example for creating new VM through client
	// newVm, err := client.CreateVM(minicloud.VM{
	// 	Name:   "demo-a",
	// 	Image:  "nginx:latest",
	// 	CPU:    0.5,
	// 	Memory: 64,
	// 	Ports:  []int{80, 443},
	// 	Env: map[string]string{
	// 		"ENV":    "dev",
	// 		"REGION": "weu",
	// 	},
	// 	ContainerID: "manual-cnt-id-for-test",
	// })

	// if err != nil {
	// 	log.Fatalf("Failed to creat VM: %v", err.Error())
	// }
	// fmt.Println("Created new VM:", newVm)

	// // Example for updating an existing VM through client
	// vmTargetID := int64(15)
	// updatedVm, err := client.UpdateVM(vmTargetID, minicloud.VM{
	// 	Name:   "demo-a",
	// 	Image:  "nginx:latest",
	// 	CPU:    0.5,
	// 	Memory: 64,
	// 	Ports:  []int{8080, 8443},
	// 	Env: map[string]string{

	// 		"ENV": "staging",
	// 	},
	// })

	// if err != nil {
	// 	log.Fatalf("Failed to update the VM: %v", err.Error())
	// }
	// fmt.Println("Udated the VM:", updatedVm)
}
