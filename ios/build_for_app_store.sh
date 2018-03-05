#!/bin/bash
set -e

pushd "$(dirname "$0")" > /dev/null

# build the actual app binary with buck
buck build ios:QCastPackage#iphoneos-arm64

# resign the app with the right entitlements because buck screws it up
fastlane sigh resign ../buck-out/gen/ios/QCastPackage\#iphoneos-arm64.ipa --signing_identity "iPhone Distribution: Stephen Magro (2DAZ3V2PJG)" -p "ProvisioningProfiles/QCast_App_Store.mobileprovision" --entitlements QCast/Entitlements.plist

popd > /dev/null
