package cmd

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/clscli/clscli/internal/cls"
	"github.com/clscli/clscli/internal/output"
	"github.com/spf13/cobra"
)

var (
	contextTopic string
	contextBTime string
	contextPrev  int
	contextNext  int
	contextQuery string
)

var contextCmd = &cobra.Command{
	Use:   "context [PkgId] [PkgLogId]",
	Short: "Get log context",
	Long:  "Retrieve context logs around a given log (PkgId and PkgLogId from SearchLog results).",
	Args:  cobra.ExactArgs(2),
	RunE:  runContext,
}

func init() {
	rootCmd.AddCommand(contextCmd)
	contextCmd.Flags().StringVarP(&contextTopic, "topic", "t", "", "Topic ID (required)")
	contextCmd.MarkFlagRequired("topic")
	contextCmd.Flags().StringVar(&contextBTime, "btime", "", "Log time (required). Accepts search result Time Unix ms or format: YYYY-mm-dd HH:MM:SS.FFF")
	contextCmd.MarkFlagRequired("btime")
	contextCmd.Flags().IntVar(&contextPrev, "prev", 0, "Number of preceding logs to retrieve (default 10 by API)")
	contextCmd.Flags().IntVar(&contextNext, "next", 0, "Number of following logs to retrieve (default 10 by API)")
	contextCmd.Flags().StringVarP(&contextQuery, "query", "q", "", "Filter context logs with query condition (no SQL)")
}

func runContext(cmd *cobra.Command, args []string) error {
	client, err := getCLSClient()
	if err != nil {
		return err
	}
	pkgID := args[0]
	pkgLogIDStr := args[1]
	pkgLogID, err := cls.PkgLogIdFromString(pkgLogIDStr)
	if err != nil {
		return fmt.Errorf("invalid PkgLogId %q: %w", pkgLogIDStr, err)
	}
	btime, err := normalizeBTime(contextBTime)
	if err != nil {
		return err
	}

	f, p := resolveOutput(cmd)
	writer, err := output.NewWriter(f, p)
	if err != nil {
		return err
	}
	defer writer.Close()

	in := cls.GetContextInput{
		TopicId:  contextTopic,
		PkgId:    pkgID,
		PkgLogId: pkgLogID,
		BTime:    btime,
		PrevLogs: int64(contextPrev),
		NextLogs: int64(contextNext),
		Query:    contextQuery,
	}
	logs, err := client.GetContext(context.Background(), in)
	if err != nil {
		return err
	}
	_ = writer.WriteTableHeaderLogContext()
	for _, log := range logs {
		if err := writer.WriteLogContextInfo(log); err != nil {
			return err
		}
	}
	return writer.Flush()
}

func normalizeBTime(value string) (string, error) {
	if value == "" {
		return "", fmt.Errorf("set --btime from the query result Time field")
	}
	if ms, err := strconv.ParseInt(value, 10, 64); err == nil {
		return time.UnixMilli(ms).Format("2006-01-02 15:04:05.000"), nil
	}
	return value, nil
}
