package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bitmovin/bitmovin-go/bitmovin"
	// "github.com/bitmovin/bitmovin-go/bitmovintypes"
	// "github.com/bitmovin/bitmovin-go/models"
	// "github.com/bitmovin/bitmovin-go/services"
	"../../../bitmovintypes"
	"../../../models"
	"../../../services"
)

func main() {
	// Creating Bitmovin object
	bitmovin := bitmovin.NewBitmovin("2f44b2a4-3f93-49d9-9408-0b51c39df667", "https://api.bitmovin.com/v1/", 5)

	prewarmedEncoderPoolS := services.NewPrewarmedEncoderPoolService(bitmovin)

	if prewarmedEncoderPools, err := prewarmedEncoderPoolS.List(0, 10); err == nil {
		fmt.Println("before creating pools:")
		for i, pool := range *prewarmedEncoderPools {
			fmt.Printf("%d: %v\n", i, pool)
		}
	} else {
		errorHandler(err)
	}

	fmt.Println("create PrewarmedEncoderPool")
	prewarmedEncoderPool := &models.CreatePrewarmedEncoderPoolRequest{
		Name:           stringToPtr("rdb PrewarmedEncoderPool test"),
		EncoderVersion: bitmovintypes.EncoderVersion("2.48.0"),
		CloudRegion:    bitmovintypes.CloudRegionAWSUSEast1,
		DiskSize:       bitmovintypes.DiskSize500GB,
		TargetPoolSize: 2,
	}
	prewarmedEncoderPoolResp, err := prewarmedEncoderPoolS.Create(prewarmedEncoderPool)
	errorHandler(err)

	if prewarmedEncoderPools, err := prewarmedEncoderPoolS.List(0, 10); err == nil {
		fmt.Println("after creating pools:")
		for i, pool := range *prewarmedEncoderPools {
			fmt.Printf("%d: %v\n", i, pool)
		}
	} else {
		errorHandler(err)
	}

	fmt.Println("start pool: " + prewarmedEncoderPoolResp.Data.Result.ID)
	_, err = prewarmedEncoderPoolS.Start(prewarmedEncoderPoolResp.Data.Result.ID)
	errorHandler(err)
	fmt.Println("wait for status: " + "STARTED")
	waitForStatus(prewarmedEncoderPoolS, prewarmedEncoderPoolResp.Data.Result.ID, "STARTED")

	fmt.Println("stop pool: " + prewarmedEncoderPoolResp.Data.Result.ID)
	_, err = prewarmedEncoderPoolS.Stop(prewarmedEncoderPoolResp.Data.Result.ID)
	errorHandler(err)
	fmt.Println("wait for status: " + "STOPPED")
	waitForStatus(prewarmedEncoderPoolS, prewarmedEncoderPoolResp.Data.Result.ID, "STOPPED")

	if _, err := prewarmedEncoderPoolS.Delete(prewarmedEncoderPoolResp.Data.Result.ID); err != nil {
		errorHandler(err)
	}

	if prewarmedEncoderPools, err := prewarmedEncoderPoolS.List(0, 10); err == nil {
		fmt.Println("after deleting pools:")
		for i, pool := range *prewarmedEncoderPools {
			fmt.Printf("%d: %v\n", i, pool)
		}
	} else {
		errorHandler(err)
	}
}

func waitForStatus(service *services.PrewarmedEncoderPoolService, poolId string, status string) {
	var st string = ""
	for st != status {
		time.Sleep(3 * time.Second)
		statusResp, err := service.Retrieve(poolId)
		if err != nil {
			fmt.Println("error in Encoding Status")
			fmt.Println(err)
			os.Exit(1)
		}
		// Polling and Printing out the response
		st = statusResp.Data.Result.Status
		fmt.Printf("status: %s\n", st)
	}
}

func errorHandler(err error) {
	if err != nil {
		switch err.(type) {
		case models.BitmovinError:
			fmt.Println("Bitmovin Error")
		default:
			fmt.Println("General Error")
		}
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func stringToPtr(s string) *string {
	return &s
}

func intToPtr(i int64) *int64 {
	return &i
}

func boolToPtr(b bool) *bool {
	return &b
}

func floatToPtr(f float64) *float64 {
	return &f
}
