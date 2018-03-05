/**
 * Copyright (c) 2015-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#import <FBSimulatorControl/FBSimulator.h>

NS_ASSUME_NONNULL_BEGIN

@class FBSimulatorApplicationOperation;
@class FBApplicationLaunchConfiguration;
@class FBDiagnostic;
@class FBProcessInfo;
@class FBSimulator;
@class FBSimulatorAgentOperation;
@class FBSimulatorConnection;
@class FBTestManager;
@protocol FBTerminationHandle;
@protocol FBJSONSerializable;

/**
 A reciever of Simulator Events
 */
@protocol FBSimulatorEventSink <NSObject>

/**
 Event for the launch of a Simulator's Container Application Process.
 This is the Simulator.app's Process.

 @param applicationProcess the Process Information for the launched Application Process.
 */
- (void)containerApplicationDidLaunch:(FBProcessInfo *)applicationProcess;

/**
 Event for the launch of a Simulator's Container Application Process.
 This is the Simulator.app's Process.

 @param applicationProcess the Process Information for the terminated Application Process.
 @param expected whether the termination was expected or not.
 */
- (void)containerApplicationDidTerminate:(FBProcessInfo *)applicationProcess expected:(BOOL)expected;

/**
 Event for the Direct Launch of a Simulator Bridge.

 @param connection the Simulator Bridge of the Simulator.
 */
- (void)connectionDidConnect:(FBSimulatorConnection *)connection;

/**
 Event for the termination of a Simulator Framebuffer.

 @param connection the Simulator Bridge of the Simulator.
 @param expected whether the termination was expected or not.
 */
- (void)connectionDidDisconnect:(FBSimulatorConnection *)connection expected:(BOOL)expected;

/**
 Event for the launch of a Simulator's launchd_sim.

 @param launchdProcess the launchd_sim process
 */
- (void)simulatorDidLaunch:(FBProcessInfo *)launchdProcess;

/**
 Event for the termination of a Simulator's launchd_sim.

 @param launchdProcess the launchd_sim process
 */
- (void)simulatorDidTerminate:(FBProcessInfo *)launchdProcess expected:(BOOL)expected;

/**
 Event for the launch of an Agent.

 @param operation the Launched Agent Operation.
 */
- (void)agentDidLaunch:(FBSimulatorAgentOperation *)operation;

/**
 Event of the termination of an agent.

 @param operation the Terminated. Agent Operation.
 @param statLoc the termination status. Documented in waitpid(2).
 */
- (void)agentDidTerminate:(FBSimulatorAgentOperation *)operation statLoc:(int)statLoc;

/**
 Event for the launch of an Application.

 @param operation the Application Operation.
 */
- (void)applicationDidLaunch:(FBSimulatorApplicationOperation *)operation;

/**
 Event for the termination of an Application.

 @param operation the Application Operation.
 @param expected whether the termination was expected or not.
 */
- (void)applicationDidTerminate:(FBSimulatorApplicationOperation *)operation expected:(BOOL)expected;

/**
 Event for the availablilty of a new log.

 @param diagnostic the diagnostic log.
 */
- (void)diagnosticAvailable:(FBDiagnostic *)diagnostic;

/**
 Event for the change in a Simulator's state.

 @param state the changed state.
 */
- (void)didChangeState:(FBSimulatorState)state;

/**
 Event for the availibility of new Termination Handle.

 @param terminationHandle the Termination Handle that is required to be called on Simulator teardown.
 */
- (void)terminationHandleAvailable:(id<FBTerminationHandle>)terminationHandle;

@end

NS_ASSUME_NONNULL_END
