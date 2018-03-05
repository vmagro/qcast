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

@class FBSimulator;
@class FBSimulatorConfiguration;
@class FBSimulatorControl;
@class FBSimulatorControlConfiguration;
@class FBSimulatorProcessFetcher;
@class FBiOSTargetQuery;
@class SimDeviceSet;

@protocol FBControlCoreLogger;

NS_ASSUME_NONNULL_BEGIN

/**
 Complements SimDeviceSet with additional functionality and more resiliant behaviours.
 Performs the preconditions necessary to call certain SimDeviceSet/SimDevice methods.
 */
@interface FBSimulatorSet : NSObject <FBDebugDescribeable, FBJSONSerializable>

#pragma mark Intializers

/**
 Creates and returns an FBSimulatorSet, performing the preconditions defined in the configuration.

 @param configuration the configuration to use. Must not be nil.
 @param deviceSet the Device Set to wrap.
 @param logger the logger to use to verbosely describe what is going on. May be nil.
 @param error any error that occurred during the creation of the pool.
 @return a new FBSimulatorPool.
 */
+ (instancetype)setWithConfiguration:(FBSimulatorControlConfiguration *)configuration deviceSet:(SimDeviceSet *)deviceSet logger:(nullable id<FBControlCoreLogger>)logger error:(NSError **)error;

#pragma mark Querying

/**
 Fetches the Simulators from the Set, matching the query.

 @param query the Query to query with.
 @return an array of matching Simulators.
 */
- (NSArray<FBSimulator *> *)query:(FBiOSTargetQuery *)query;

#pragma mark Creation Methods

/**
 Creates and returns a FBSimulator fbased on a configuration.

 @param configuration the Configuration of the Device to Allocate. Must not be nil.
 @return a Future wrapping a created FBSimulator if one could be allocated with the provided options.
 */
- (FBFuture<FBSimulator *> *)createSimulatorWithConfiguration:(FBSimulatorConfiguration *)configuration;

/**
 Finds and creates the Configurations for the missing 'Default Simulators' in the reciever.
 */
- (NSArray<FBSimulatorConfiguration *> *)configurationsForAbsentDefaultSimulators;

#pragma mark Desctructive Methods

/**
 Kills a Simulator in the Set.
 The Set to which the Simulator belongs must be the reciever.

 @param simulator the Simulator to delete. Must not be nil.
 @return an Future that resolves when successful.
 */
- (FBFuture<NSArray<FBSimulator *> *> *)killSimulator:(FBSimulator *)simulator;

/**
 Erases a Simulator in the Set.
 The Set to which the Simulator belongs must be the reciever.

 @param simulator the Simulator to erase. Must not be nil.
 @return A future wrapping the erased simulators udids.
 */
- (FBFuture<NSArray<FBSimulator *> *> *)eraseSimulator:(FBSimulator *)simulator;

/**
 Deletes a Simulator in the Set.
 The Set to which the Simulator belongs must be the reciever.

 @param simulator the Simulator to delete. Must not be nil.
 @return A future wrapping the delegate simulators.
 */
- (FBFuture<NSArray<NSString *> *> *)deleteSimulator:(FBSimulator *)simulator;

/**
 Kills all provided Simulators.
 The Set to which the Simulators belong must be the reciever.

 @param simulators the Simulators to kill. Must not be nil.
 @return an Future that resolves when successful.
 */
- (FBFuture<NSArray<FBSimulator *> *> *)killAll:(NSArray<FBSimulator *> *)simulators;

/**
 Erases all provided Simulators.
 The Set to which the Simulators belong must be the reciever.

 @param simulators the Simulators to erase. Must not be nil.
 @return A future wrapping the erased simulators.
 */
- (FBFuture<NSArray<FBSimulator *> *> *)eraseAll:(NSArray<FBSimulator *> *)simulators;

/**
 Erases all provided Simulators.
 The Set to which the Simulators belong must be the reciever.

 @param simulators the Simulators to delete. Must not be nil.
 @return A future wrapping the erased simulators udids.
 */
- (FBFuture<NSArray<NSString *> *> *)deleteAll:(NSArray<FBSimulator *> *)simulators;

/**
 Kills all of the Simulators that belong to the reciever.

 @return an Future that resolves when successful.
 */
- (FBFuture<NSArray<FBSimulator *> *> *)killAll;

/**
 Kills all of the Simulators that belong to the reciever.

 @return A future wrapping the erased simulators.
 */
- (FBFuture<NSArray<FBSimulator *> *> *)eraseAll;

/**
 Delete all of the Simulators that belong to the reciever.

 @return A future wrapping the erased simulators udids.
 */
- (FBFuture<NSArray<NSString *> *> *)deleteAll;

/**
 The Logger to use.
 */
@property (nonatomic, strong, readonly) id<FBControlCoreLogger> logger;

/**
 Returns the configuration for the reciever.
 */
@property (nonatomic, copy, readonly) FBSimulatorControlConfiguration *configuration;

/**
 The SimDeviceSet to that is owned by the reciever.
 */
@property (nonatomic, strong, readonly) SimDeviceSet *deviceSet;

/**
 The FBProcessFetcher that is used to obtain Simulator Process Information.
 */
@property (nonatomic, strong, readonly) FBSimulatorProcessFetcher *processFetcher;

/**
 An NSArray<FBSimulator> of all Simulators in the Set.
*/
@property (nonatomic, copy, readonly) NSArray<FBSimulator *> *allSimulators;

@end

NS_ASSUME_NONNULL_END
