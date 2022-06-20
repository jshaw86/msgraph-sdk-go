package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidGeneralDeviceConfigurationable 
type AndroidGeneralDeviceConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppsBlockClipboardSharing()(*bool)
    GetAppsBlockCopyPaste()(*bool)
    GetAppsBlockYouTube()(*bool)
    GetAppsHideList()([]AppListItemable)
    GetAppsInstallAllowList()([]AppListItemable)
    GetAppsLaunchBlockList()([]AppListItemable)
    GetBluetoothBlocked()(*bool)
    GetCameraBlocked()(*bool)
    GetCellularBlockDataRoaming()(*bool)
    GetCellularBlockMessaging()(*bool)
    GetCellularBlockVoiceRoaming()(*bool)
    GetCellularBlockWiFiTethering()(*bool)
    GetCompliantAppListType()(*AppListType)
    GetCompliantAppsList()([]AppListItemable)
    GetDeviceSharingAllowed()(*bool)
    GetDiagnosticDataBlockSubmission()(*bool)
    GetFactoryResetBlocked()(*bool)
    GetGoogleAccountBlockAutoSync()(*bool)
    GetGooglePlayStoreBlocked()(*bool)
    GetKioskModeApps()([]AppListItemable)
    GetKioskModeBlockSleepButton()(*bool)
    GetKioskModeBlockVolumeButtons()(*bool)
    GetLocationServicesBlocked()(*bool)
    GetNfcBlocked()(*bool)
    GetPasswordBlockFingerprintUnlock()(*bool)
    GetPasswordBlockTrustAgents()(*bool)
    GetPasswordExpirationDays()(*int32)
    GetPasswordMinimumLength()(*int32)
    GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32)
    GetPasswordPreviousPasswordBlockCount()(*int32)
    GetPasswordRequired()(*bool)
    GetPasswordRequiredType()(*AndroidRequiredPasswordType)
    GetPasswordSignInFailureCountBeforeFactoryReset()(*int32)
    GetPowerOffBlocked()(*bool)
    GetScreenCaptureBlocked()(*bool)
    GetSecurityRequireVerifyApps()(*bool)
    GetStorageBlockGoogleBackup()(*bool)
    GetStorageBlockRemovableStorage()(*bool)
    GetStorageRequireDeviceEncryption()(*bool)
    GetStorageRequireRemovableStorageEncryption()(*bool)
    GetVoiceAssistantBlocked()(*bool)
    GetVoiceDialingBlocked()(*bool)
    GetWebBrowserBlockAutofill()(*bool)
    GetWebBrowserBlocked()(*bool)
    GetWebBrowserBlockJavaScript()(*bool)
    GetWebBrowserBlockPopups()(*bool)
    GetWebBrowserCookieSettings()(*WebBrowserCookieSettings)
    GetWiFiBlocked()(*bool)
    SetAppsBlockClipboardSharing(value *bool)()
    SetAppsBlockCopyPaste(value *bool)()
    SetAppsBlockYouTube(value *bool)()
    SetAppsHideList(value []AppListItemable)()
    SetAppsInstallAllowList(value []AppListItemable)()
    SetAppsLaunchBlockList(value []AppListItemable)()
    SetBluetoothBlocked(value *bool)()
    SetCameraBlocked(value *bool)()
    SetCellularBlockDataRoaming(value *bool)()
    SetCellularBlockMessaging(value *bool)()
    SetCellularBlockVoiceRoaming(value *bool)()
    SetCellularBlockWiFiTethering(value *bool)()
    SetCompliantAppListType(value *AppListType)()
    SetCompliantAppsList(value []AppListItemable)()
    SetDeviceSharingAllowed(value *bool)()
    SetDiagnosticDataBlockSubmission(value *bool)()
    SetFactoryResetBlocked(value *bool)()
    SetGoogleAccountBlockAutoSync(value *bool)()
    SetGooglePlayStoreBlocked(value *bool)()
    SetKioskModeApps(value []AppListItemable)()
    SetKioskModeBlockSleepButton(value *bool)()
    SetKioskModeBlockVolumeButtons(value *bool)()
    SetLocationServicesBlocked(value *bool)()
    SetNfcBlocked(value *bool)()
    SetPasswordBlockFingerprintUnlock(value *bool)()
    SetPasswordBlockTrustAgents(value *bool)()
    SetPasswordExpirationDays(value *int32)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)()
    SetPasswordPreviousPasswordBlockCount(value *int32)()
    SetPasswordRequired(value *bool)()
    SetPasswordRequiredType(value *AndroidRequiredPasswordType)()
    SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)()
    SetPowerOffBlocked(value *bool)()
    SetScreenCaptureBlocked(value *bool)()
    SetSecurityRequireVerifyApps(value *bool)()
    SetStorageBlockGoogleBackup(value *bool)()
    SetStorageBlockRemovableStorage(value *bool)()
    SetStorageRequireDeviceEncryption(value *bool)()
    SetStorageRequireRemovableStorageEncryption(value *bool)()
    SetVoiceAssistantBlocked(value *bool)()
    SetVoiceDialingBlocked(value *bool)()
    SetWebBrowserBlockAutofill(value *bool)()
    SetWebBrowserBlocked(value *bool)()
    SetWebBrowserBlockJavaScript(value *bool)()
    SetWebBrowserBlockPopups(value *bool)()
    SetWebBrowserCookieSettings(value *WebBrowserCookieSettings)()
    SetWiFiBlocked(value *bool)()
}
