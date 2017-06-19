// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: github.com/overcyn/matcha/pb/view/stacknav/stacknavigator.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers.h>
#else
 #import "GPBProtocolBuffers.h"
#endif

#if GOOGLE_PROTOBUF_OBJC_VERSION < 30002
#error This file was generated by a newer version of protoc which is incompatible with your Protocol Buffer library sources.
#endif
#if 30002 < GOOGLE_PROTOBUF_OBJC_MIN_SUPPORTED_VERSION
#error This file was generated by an older version of protoc which is incompatible with your Protocol Buffer library sources.
#endif

// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

CF_EXTERN_C_BEGIN

@class MatchaPBStackNavScreen;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - MatchaPBStackNavStacknavigatorRoot

/**
 * Exposes the extension registry for this file.
 *
 * The base class provides:
 * @code
 *   + (GPBExtensionRegistry *)extensionRegistry;
 * @endcode
 * which is a @c GPBExtensionRegistry that includes all the extensions defined by
 * this file and all files that it depends on.
 **/
@interface MatchaPBStackNavStacknavigatorRoot : GPBRootObject
@end

#pragma mark - MatchaPBStackNavScreen

typedef GPB_ENUM(MatchaPBStackNavScreen_FieldNumber) {
  MatchaPBStackNavScreen_FieldNumber_Id_p = 1,
  MatchaPBStackNavScreen_FieldNumber_Title = 2,
  MatchaPBStackNavScreen_FieldNumber_BackButtonTitle = 3,
  MatchaPBStackNavScreen_FieldNumber_BackButtonHidden = 4,
  MatchaPBStackNavScreen_FieldNumber_TitleViewId = 5,
  MatchaPBStackNavScreen_FieldNumber_RightViewIdsArray = 6,
  MatchaPBStackNavScreen_FieldNumber_LeftViewIdsArray = 7,
  MatchaPBStackNavScreen_FieldNumber_CustomBackButtonTitle = 8,
};

@interface MatchaPBStackNavScreen : GPBMessage

@property(nonatomic, readwrite) int64_t id_p;

@property(nonatomic, readwrite, copy, null_resettable) NSString *title;

@property(nonatomic, readwrite) BOOL customBackButtonTitle;

@property(nonatomic, readwrite, copy, null_resettable) NSString *backButtonTitle;

@property(nonatomic, readwrite) BOOL backButtonHidden;

@property(nonatomic, readwrite) int64_t titleViewId;

@property(nonatomic, readwrite, strong, null_resettable) GPBInt64Array *rightViewIdsArray;
/** The number of items in @c rightViewIdsArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger rightViewIdsArray_Count;

@property(nonatomic, readwrite, strong, null_resettable) GPBInt64Array *leftViewIdsArray;
/** The number of items in @c leftViewIdsArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger leftViewIdsArray_Count;

@end

#pragma mark - MatchaPBStackNavStackNav

typedef GPB_ENUM(MatchaPBStackNavStackNav_FieldNumber) {
  MatchaPBStackNavStackNav_FieldNumber_ScreensArray = 1,
};

@interface MatchaPBStackNavStackNav : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<MatchaPBStackNavScreen*> *screensArray;
/** The number of items in @c screensArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger screensArray_Count;

@end

#pragma mark - MatchaPBStackNavStackEvent

typedef GPB_ENUM(MatchaPBStackNavStackEvent_FieldNumber) {
  MatchaPBStackNavStackEvent_FieldNumber_IdArray = 1,
};

@interface MatchaPBStackNavStackEvent : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) GPBInt64Array *idArray;
/** The number of items in @c idArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger idArray_Count;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
