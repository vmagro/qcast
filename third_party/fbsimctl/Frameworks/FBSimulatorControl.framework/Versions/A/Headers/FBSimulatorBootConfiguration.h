/**
 * Copyright (c) 2015-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#import <Foundation/Foundation.h>

#import <FBControlCore/FBControlCore.h>

@class FBFramebufferConfiguration;

/**
 An Option Set for Direct Launching.
 */
typedef NS_OPTIONS(NSUInteger, FBSimulatorBootOptions) {
  FBSimulatorBootOptionsConnectBridge = 1 << 0, /** Connects the Simulator Bridge on boot, rather than lazily on-demand */
  FBSimulatorBootOptionsEnableDirectLaunch = 1 << 1, /** Launches the Simulator via directly (via SimDevice) instead of with Simulator.app. Enables Framebuffer Connection. */
  FBSimulatorBootOptionsUseNSWorkspace = 1 << 2, /** Uses -[NSWorkspace launchApplicationAtURL:options:configuration::error:] to launch Simulator.app */
  FBSimulatorBootOptionsAwaitServices = 1 << 3, /** Waits for the Simulator to be in a 'Usable' state before returning on the boot command */
};

NS_ASSUME_NONNULL_BEGIN

/**
 The Action Type for a Boot.
 */
extern FBiOSTargetActionType const FBiOSTargetActionTypeBoot;

/**
 A Value Object for defining how to launch a Simulator.
 */
@interface FBSimulatorBootConfiguration : NSObject <NSCopying, FBiOSTargetFuture, FBDebugDescribeable>

/**
 Options for how the Simulator should be launched.
 */
@property (nonatomic, assign, readonly) FBSimulatorBootOptions options;

/**
 The Locale in which to Simulate, may be nil.
 */
@property (nonatomic, nullable, strong, readonly) FBLocalizationOverride *localizationOverride;

/**
 The Scale of the Framebuffer.
 */
@property (nonatomic, nullable, copy, readonly) FBScale scale;

/**
 Configuration for the Framebuffer.
 If nil, means that the Framebuffer will not be connected on launch
 */
@property (nonatomic, nullable, copy, readonly) FBFramebufferConfiguration *framebuffer;

#pragma mark Default Instance

/**
 The Default Configuration.
 */
@property (nonatomic, strong, class, readonly) FBSimulatorBootConfiguration *defaultConfiguration;

#pragma mark Launch Options

/**
 Set Direct Launch Options
 */
+ (instancetype)withOptions:(FBSimulatorBootOptions)options;
- (instancetype)withOptions:(FBSimulatorBootOptions)options;

#pragma mark Device Scale

/**
 Returns a new Configuration with the Scale Applied.
 */
+ (instancetype)withScale:(nullable FBScale)scale;
- (instancetype)withScale:(nullable FBScale)scale;

#pragma mark Locale

/**
 Set the Localization Override
 */
+ (instancetype)withLocalizationOverride:(nullable FBLocalizationOverride *)localizationOverride;
- (instancetype)withLocalizationOverride:(nullable FBLocalizationOverride *)localizationOverride;

#pragma mark Framebuffer

/**
 Set Framebuffer Configuration
 */
+ (instancetype)withFramebuffer:(nullable FBFramebufferConfiguration *)framebuffer;
- (instancetype)withFramebuffer:(nullable FBFramebufferConfiguration *)framebuffer;

@end

NS_ASSUME_NONNULL_END
