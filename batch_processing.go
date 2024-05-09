package main

import (
    "context"
    "fmt"
    "google.golang.org/api/dataproc/v1"
    "google.golang.org/api/option"
)

func main() {
    // Initialize a context
    ctx := context.Background()

    // Create a Dataproc client
    client, err := dataproc.NewService(ctx, option.WithCredentialsFile("/path/to/your/service-account-key.json"))
    if err != nil {
        fmt.Println("Failed to create Dataproc client:", err)
        return
    }

    // Define the cluster configuration
    cluster := &dataproc.Cluster{
        Config: &dataproc.ClusterConfig{
            MasterConfig: &dataproc.InstanceGroupConfig{
                NumInstances:   1,
                MachineTypeUri: "n1-standard-4",
            },
            WorkerConfig: &dataproc.InstanceGroupConfig{
                NumInstances:   2,
                MachineTypeUri: "n1-standard-4",
            },
            ConfigBucket: "gs://your-bucket/",
        },
    }

    // Set the project and region
    project := "your-project-id"
    region := "your-region"

    // Create a cluster
    createOp, err := client.Projects.Regions.Clusters.Create(project, region, cluster).Do()
    if err != nil {
        fmt.Println("Failed to create cluster:", err)
        return
    }

    // Get operation status
    statusOp, err := client.Projects.Regions.Operations.Get(project, region, createOp.Name).Do()
    if err != nil {
        fmt.Println("Failed to get operation status:", err)
        return
    }

    fmt.Println("Cluster created successfully:", statusOp.Name)
}
