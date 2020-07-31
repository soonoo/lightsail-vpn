package main

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/lightsail"
)

func main() {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-2"),
        Credentials: credentials.NewSharedCredentials("", "lightsail-vpn"),
    })
    if err != nil {
        fmt.Println(err)
        return
    }

    // svc := lightsail.New(sess)
    // result, err := svc.GetInstances(&lightsail.GetInstancesInput{})
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // fmt.Println(result)

    // svc := lightsail.New(sess)
    // result, err := svc.GetInstancePortStates(&lightsail.GetInstancePortStatesInput{
    //     InstanceName: aws.String("lightsail-api-test"),
    // })
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // fmt.Println(result)

    // res, err := svc.GetBlueprints(&lightsail.GetBlueprintsInput{})
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // fmt.Println(res)

    // res, err := svc.GetBundles(&lightsail.GetBundlesInput{})
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // fmt.Println(res)

    // res, err := svc.CreateInstances(&lightsail.CreateInstancesInput{
    //     AvailabilityZone: aws.String("us-west-2a"),
    //     BlueprintId: aws.String("ubuntu_16_04_2"),
    //     BundleId: aws.String("nano_2_0"),
    //     InstanceNames: aws.StringSlice([]string{"lightsail-api-test"}),
    // })
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // fmt.Println(res)

    // result, err := svc.DownloadDefaultKeyPair(&lightsail.DownloadDefaultKeyPairInput{})
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // fmt.Println(result)
}
