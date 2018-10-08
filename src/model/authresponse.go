package model

type AuthResponse struct{
	Status string `json:"status"`
	DeveleperPermissionDevices string `json:"develeperPermissionDevices"`
	RemotePermissionDevices string `json:"remotePermissionDevices"`
}

