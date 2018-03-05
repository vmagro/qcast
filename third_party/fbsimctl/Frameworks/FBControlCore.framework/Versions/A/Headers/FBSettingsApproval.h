/**
 * Copyright (c) 2015-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#import <Foundation/Foundation.h>

#import <FBControlCore/FBJSONConversion.h>

NS_ASSUME_NONNULL_BEGIN

/**
 An Enumeration Representing Approval of Services.
 */
typedef NSString *FBSettingsApprovalService NS_STRING_ENUM;

extern FBSettingsApprovalService const FBSettingsApprovalServiceContacts;
extern FBSettingsApprovalService const FBSettingsApprovalServicePhotos;
extern FBSettingsApprovalService const FBSettingsApprovalServiceCamera;
extern FBSettingsApprovalService const FBSettingsApprovalServiceLocation;
extern FBSettingsApprovalService const FBSettingsApprovalServiceMicrophone;

/**
 Value container for approval of settings.
 */
@interface FBSettingsApproval : NSObject <NSCopying, FBJSONSerializable, FBJSONDeserializable>

#pragma mark Initializers

/**
 The Designated Initializer

 @param bundleIDs the bundle ids to apply the approvals to.
 @param services the services that will be approved.
 @return a new FBSettingsApproval Instance.
 */
+ (instancetype)approvalWithBundleIDs:(NSArray<NSString *> *)bundleIDs services:(NSArray<FBSettingsApprovalService> *)services;

#pragma mark Properties

/**
 The Bundle IDs to apply the approval to.
 */
@property (nonatomic, copy, readonly) NSArray<NSString *> *bundleIDs;

/**
 The Bundle IDs to apply the approval to.
 */
@property (nonatomic, copy, readonly) NSArray<FBSettingsApprovalService> *services;

@end

NS_ASSUME_NONNULL_END
