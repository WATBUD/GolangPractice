// Code generated by goa v3.16.2, DO NOT EDIT.
//
// HTTP request path constructors for the Sync service.
//
// Command:
// $ goa gen mai.today/api/design

package server

import (
	"fmt"
)

// ReceiveSyncSyncPath returns the URL path to the Sync service receiveSync HTTP endpoint.
func ReceiveSyncSyncPath(channel string) string {
	return fmt.Sprintf("/%v/sync", channel)
}

// SyncSyncPath returns the URL path to the Sync service sync HTTP endpoint.
func SyncSyncPath() string {
	return "/sync"
}