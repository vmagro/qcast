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

NS_ASSUME_NONNULL_BEGIN

@class FBProcessInfo;
@class FBSimulator;

/**
 Protocol for interacting with a Simulator's launchctl
 */
@protocol FBSimulatorLaunchCtlCommands <NSObject, FBiOSTargetCommand>

#pragma mark Processes

/*
 Fetches an NSArray<FBProcessInfo *> of the subprocesses of the launchd_sim.
 */
- (NSArray<FBProcessInfo *> *)launchdSimSubprocesses;

#pragma mark Querying Services

/**
 Finds the Service Name for a provided process.
 Will fail if there is no process matching the Process Info found.

 @param process the process to obtain the name for.
 @return A Future, wrapping the Service Name.
 */
- (FBFuture<NSString *> *)serviceNameForProcess:(FBProcessInfo *)process;

/**
 Finds the Service Name for a given Service Name.
 This will perform a substring match, so the first matching service will be returned.

 @param serviceName a Substring of the Service to fetch.
 @return A Future, wrapping a tuple of String Service Name & NSNumber Process Identifier.
 */
- (FBFuture<NSArray<id> *> *)serviceNameAndProcessIdentifierForSubstring:(NSString *)serviceName;

/**
 Consults the Simulator's launchctl to determine if the given process.

 @param process the process to look for.
 @return A Future, YES if the process exists. NO otherwise.
 */
- (FBFuture<NSNumber *> *)processIsRunningOnSimulator:(FBProcessInfo *)process;

/**
 Returns the currently running launchctl services.
 Returns a Mapping of Service Name to Process Identifier.
 NSNull is used to represent services that do not have a Process Identifier.

 @return A Future, wrapping a Mapping of Service Name to Process identifier.
 */
- (FBFuture<NSDictionary<NSString *, id> *> *)listServices;

#pragma mark Manipulating Services

/**
 Stops the Provided Process, by Service Name.

 @param serviceName the name of the Process to Stop.
 @return A Future, wrapping the Service Name of the Stopped process, or nil if the process does not exist.
 */
- (FBFuture<NSString *> *)stopServiceWithName:(NSString *)serviceName;

/**
 Starts the Provided Process, by Service Name.

 @param serviceName the name of the Process to Stop.
 @return A Future, wrapping the Service Name of the Stopped process, or nil if the process does not exist.
 */
- (FBFuture<NSString *> *)startServiceWithName:(NSString *)serviceName;

@end

/**
 An Interface to a Simulator's launchctl.
 */
@interface FBSimulatorLaunchCtlCommands : NSObject <FBSimulatorLaunchCtlCommands>

@end

NS_ASSUME_NONNULL_END
