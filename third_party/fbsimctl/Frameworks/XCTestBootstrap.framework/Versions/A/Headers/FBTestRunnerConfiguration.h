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

@class FBApplicationDataPackage;
@class FBProductBundle;
@class FBTestBundle;
@class FBTestConfiguration;

/**
 A Configuration Value for the Test Runner.
 */
@interface FBTestRunnerConfiguration : NSObject <NSCopying>

/**
 The Designated Initializer

 @param sessionIdentifier identifier used to run test.
 @param hostApplication the test host.
 @param ideInjectionFramework IDEBundleInjection.framework product bundle.
 @param testBundle the test bundle.
 @param testConfigurationPath path to test configuration that should be used to start tests.
 @param frameworkSearchPath the search path for Frameworks.
 @param testedApplicationAdditionalEnvironment Launch environment variables added to test target application
 */
+ (instancetype)configurationWithSessionIdentifier:(NSUUID *)sessionIdentifier hostApplication:(FBProductBundle *)hostApplication ideInjectionFramework:(FBProductBundle *)ideInjectionFramework testBundle:(FBTestBundle *)testBundle testConfigurationPath:(NSString *)testConfigurationPath frameworkSearchPath:(NSString *)frameworkSearchPath testedApplicationAdditionalEnvironment:(nullable NSDictionary<NSString *, NSString *> *)testedApplicationAdditionalEnvironment;

/**
 Test session identifier
 */
@property (nonatomic, copy, readonly) NSUUID *sessionIdentifier;

/**
 Test runner app used for testing
 */
@property (nonatomic, strong, readonly) FBProductBundle *testRunner;

/**
  Launch arguments for test runner
 */
@property (nonatomic, copy, readonly) NSArray<NSString *> *launchArguments;

/**
 Launch environment variables for test runner
 */
@property (nonatomic, copy, readonly) NSDictionary<NSString *, NSString *> *launchEnvironment;

/**
 Launch environment variables added to test target application
 */
@property (nonatomic, copy, readonly) NSDictionary<NSString *, NSString *> *testedApplicationAdditionalEnvironment;

@end

NS_ASSUME_NONNULL_END
