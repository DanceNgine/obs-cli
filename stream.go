package main

import (
	"fmt"
	
	obsws "github.com/DanceNgine/go-obs-websocket"
	"github.com/spf13/cobra"
)

var (
	startStreamCmd = &cobra.Command{
		Use:   "start-stream",
		Short: "Starts streaming",
		RunE: func(cmd *cobra.Command, args []string) error {
			return startStream()
		},
	}

	stopStreamCmd = &cobra.Command{
		Use:   "stop-stream",
		Short: "Stops streaming",
		RunE: func(cmd *cobra.Command, args []string) error {
			return stopStream()
		},
	}
	
	statusStreamCmd = &cobra.Command{
		Use:	"status-stream",
		Short: "Displays Status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return statusStream()
		},
	
	}
)

func startStream() error {
	m := make(map[string]interface{})
	req := obsws.NewStartStreamingRequest(m, "", m, m, "", "", true, "", "")
	return req.Send(*client)
}

func stopStream() error {
	req := obsws.NewStopStreamingRequest()
	return req.Send(*client)
}

func statusStream() error {
	{
		req := obsws.NewGetStreamingStatusRequest()
		resp, err := req.SendReceive(*client)
		if err != nil {
			return err
		}
		fmt.Println("Stream Status")
		fmt.Println("===============")
		fmt.Println("Streaming: ", resp.Streaming)
		fmt.Println("Time: ", resp.StreamTimecode)
	}
	
	return nil
}

func init() {
	rootCmd.AddCommand(startStreamCmd)
	rootCmd.AddCommand(stopStreamCmd)
	rootCmd.AddCommand(statusStreamCmd)
}
