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

@class FBSimulator;

/**
 Process Launch Configuration for Simulators.
 */
@interface FBProcessLaunchConfiguration (Simulator)

/**
  Adds Environment to the Launch Configuration

  @param environmentAdditions the Environment to Add. Must be an NSDictionary<NSString *, NSString*>>
  @return a new Launch Configuration with the Environment Applied.
*/
- (instancetype)withEnvironmentAdditions:(NSDictionary<NSString *, NSString *> *)environmentAdditions;

/**
 Adds Diagnostic Environment information to the reciever's environment configuration.

 @return a new Launch Configuration with the Diagnostic Environment Applied.
 */
- (instancetype)withDiagnosticEnvironment;

/**
 Uses DYLD_INSERT_LIBRARIES to inject a dylib into the launched application's process.

 @param filePath the File Path to the Dynamic Library. Must not be nil.
 */
- (instancetype)injectingLibrary:(NSString *)filePath;

/**
 Injects the Shimulator Dylib into the launched process;
 */
- (instancetype)injectingShimulator;

/**
 Creates a FBDiagnostic for the location of the stdout, if applicable.

 @param simulator the simulator to create the diagnostic for.
 @param diagnosticOut an outparam for the diagnostic.
 @param error an error out for any error that occurs.
 @return a diagnostic if applicable, nil otherwise.
 */
- (BOOL)createStdOutDiagnosticForSimulator:(FBSimulator *)simulator diagnosticOut:(FBDiagnostic *_Nullable * _Nullable)diagnosticOut error:(NSError **)error;

/**
 Creates a FBDiagnostic for the location of the stderr, if applicable.

 @param simulator the simulator to create the diagnostic for.
 @param diagnosticOut an outparam for the diagnostic.
 @param error an error out for any error that occurs.
 @return a diagnostic if applicable, nil otherwise.
 */
- (BOOL)createStdErrDiagnosticForSimulator:(FBSimulator *)simulator diagnosticOut:(FBDiagnostic *_Nullable* _Nullable)diagnosticOut error:(NSError **)error;

/**
 Creates a FBDiagnostic for the location of the selector, if applicable.

 @param selector the selector to apply.
 @param simulator the simulator to create the diagnostic for.
 @param diagnosticOut an outparam for the diagnostic.
 @param error an error out for any error that occurs.
 @return a diagnostic if applicable, nil otherwise.
 */
- (BOOL)createDiagnosticForSelector:(SEL)selector simulator:(FBSimulator *)simulator diagnosticOut:(FBDiagnostic *_Nullable*_Nullable)diagnosticOut error:(NSError **)error;

/**
 A Name used to distinguish between Launch Configurations.
 */
- (NSString *)identifiableName;

@end

/**
 Helpers for Agent Launches.
 */
@interface FBAgentLaunchConfiguration (Helpers)

/**
 Creates the Dictionary of launch options for spawning an Agent.

 @param stdOut the stdout to use, may be nil.
 @param stdErr the stderr to use, may be nil.
 @return a Dictionary if successful, nil otherwise.
 */
- (NSDictionary<NSString *, id> *)simDeviceLaunchOptionsWithStdOut:(nullable NSFileHandle *)stdOut stdErr:(nullable NSFileHandle *)stdErr;

/**
 Creates the Dictionary of launch options for spawning an Agent.
 This static method allows the options dictionary to be constructed, without an FBAgentLaunchConfiguration.

 @prarm launchPath the Launch Path.
 @param arguments the arguments.
 @param environment the environment
 @param waitForDebugger YES if the process should be launched waiting for a debugger to attach. NO otherwise.
 @param stdOut the stdout to use, may be nil.
 @param stdErr the stderr to use, may be nil.
 @return a Dictionary if successful, nil otherwise.
 */
+ (NSDictionary<NSString *, id> *)simDeviceLaunchOptionsWithLaunchPath:(NSString *)launchPath arguments:(NSArray<NSString *> *)arguments environment:(NSDictionary<NSString *, NSString *> *)environment waitForDebugger:(BOOL)waitForDebugger stdOut:(nullable NSFileHandle *)stdOut stdErr:(nullable NSFileHandle *)stdErr;

@end

/**
 Helpers for Application Launches.
 */
@interface FBApplicationLaunchConfiguration (Helpers)

/**
 Overrides the launch of the Application with a given localization.

 @param localizationOverride the Localization Override to Apply.s
 */
- (instancetype)overridingLocalization:(FBLocalizationOverride *)localizationOverride;

/**
 Creates the Dictionary of launch options for launching an Application.

 @param stdOutPath the path to launch stdout to, may be nil.
 @param stdErrPath the path to launch stderr to, may be nil.
 @param waitForDebugger YES if the Application should be launched waiting for a debugger to attach. NO otherwise.
 @return a Dictionary if successful, nil otherwise.
 */
- (NSDictionary<NSString *, id> *)simDeviceLaunchOptionsWithStdOutPath:(nullable NSString *)stdOutPath stdErrPath:(nullable NSString *)stdErrPath waitForDebugger:(BOOL)waitForDebugger;

@end

NS_ASSUME_NONNULL_END
