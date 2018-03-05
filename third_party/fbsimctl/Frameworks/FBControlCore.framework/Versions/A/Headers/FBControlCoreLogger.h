/**
 * Copyright (c) 2015-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#import <Foundation/Foundation.h>

NS_ASSUME_NONNULL_BEGIN

/**
 A Protocol for Classes that recieve Logger Messages.
 */
@protocol FBControlCoreLogger <NSObject>

/**
 Logs a Message with the provided String.

 @param string the string to log.
 @return the reciever, for chaining.
 */
- (id<FBControlCoreLogger>)log:(NSString *)string;

/**
 Logs a Message with the provided Format String.

 @param format the Format String for the Logger.
 @return the reciever, for chaining.
 */
- (id<FBControlCoreLogger>)logFormat:(NSString *)format, ... NS_FORMAT_FUNCTION(1,2);

/**
 Returns the Info Logger variant.
 */
- (id<FBControlCoreLogger>)info;

/**
 Returns the Debug Logger variant.
 */
- (id<FBControlCoreLogger>)debug;

/**
 Returns the Error Logger variant.
 */
- (id<FBControlCoreLogger>)error;

/**
 Returns a Logger that will accept log values on the given queue.

 @param queue the queue to accept log messages on.
 @return a new Logger that will allows logging of messages on the provided queue.
 */
- (id<FBControlCoreLogger>)onQueue:(dispatch_queue_t)queue;

/**
 Returns a Logger that will prefix all messages with the given string

 @param prefix the prefix to prepend to all messages.
 @return a new Logger that will allows logging of messages on the provided queue.
 */
- (id<FBControlCoreLogger>)withPrefix:(NSString *)prefix;

@end

/**
 Implementations of Loggers.
 */
@interface FBControlCoreLogger : NSObject

/**
 An implementation of `FBControlCoreLogger` that logs all events using ASL.

 @param writeToStdErr YES if all future log messages should be written to stderr, NO otherwise.
 @param debugLogging YES if Debug messages should be written to stderr, NO otherwise.
 @return an FBControlCoreLogger instance.
 */
+ (id<FBControlCoreLogger>)systemLoggerWritingToStderrr:(BOOL)writeToStdErr withDebugLogging:(BOOL)debugLogging;

/**
 An implementation of `FBControlCoreLogger` that logs all to a file descriptor using ASL.

 @param fileHandle the file handle to log to, or nil to not log to the file handle. Logging to stderr may still occur.
 @param debugLogging YES if Debug messages should be written to the file descriptor.
 @return an FBControlCoreLogger instance.
 */
+ (id<FBControlCoreLogger>)systemLoggerWritingToFileHandle:(nullable NSFileHandle *)fileHandle withDebugLogging:(BOOL)debugLogging;

/**
 Compose multiple loggers into one.
 
 @param loggers the loggers to compose.
 @return the composite logger.
 */
+ (id<FBControlCoreLogger>)compositeLoggerWithLoggers:(NSArray<id<FBControlCoreLogger>> *)loggers;

@end

NS_ASSUME_NONNULL_END
