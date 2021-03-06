// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gomatcha.io/matcha/pb/touch/touch2.proto

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

@class GPBAny;
@class GPBDuration;
@class GPBTimestamp;
@class MatchaLayoutPBPoint;
@class MatchaPBTouchRecognizer;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - Enum MatchaPBTouchEventKind

typedef GPB_ENUM(MatchaPBTouchEventKind) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  MatchaPBTouchEventKind_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  MatchaPBTouchEventKind_EventKindPossible = 0,
  MatchaPBTouchEventKind_EventKindChanged = 1,
  MatchaPBTouchEventKind_EventKindFailed = 2,
  MatchaPBTouchEventKind_EventKindRecognized = 3,
};

GPBEnumDescriptor *MatchaPBTouchEventKind_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL MatchaPBTouchEventKind_IsValidValue(int32_t value);

#pragma mark - MatchaPBTouchTouch2Root

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
@interface MatchaPBTouchTouch2Root : GPBRootObject
@end

#pragma mark - MatchaPBTouchRecognizer

typedef GPB_ENUM(MatchaPBTouchRecognizer_FieldNumber) {
  MatchaPBTouchRecognizer_FieldNumber_Id_p = 1,
  MatchaPBTouchRecognizer_FieldNumber_Recognizer = 3,
};

@interface MatchaPBTouchRecognizer : GPBMessage

@property(nonatomic, readwrite) int64_t id_p;

@property(nonatomic, readwrite, strong, null_resettable) GPBAny *recognizer;
/** Test to see if @c recognizer has been set. */
@property(nonatomic, readwrite) BOOL hasRecognizer;

@end

#pragma mark - MatchaPBTouchRecognizerList

typedef GPB_ENUM(MatchaPBTouchRecognizerList_FieldNumber) {
  MatchaPBTouchRecognizerList_FieldNumber_RecognizersArray = 1,
};

@interface MatchaPBTouchRecognizerList : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<MatchaPBTouchRecognizer*> *recognizersArray;
/** The number of items in @c recognizersArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger recognizersArray_Count;

@end

#pragma mark - MatchaPBTouchButtonRecognizer

typedef GPB_ENUM(MatchaPBTouchButtonRecognizer_FieldNumber) {
  MatchaPBTouchButtonRecognizer_FieldNumber_OnEvent = 1,
  MatchaPBTouchButtonRecognizer_FieldNumber_IgnoresScroll = 2,
};

@interface MatchaPBTouchButtonRecognizer : GPBMessage

@property(nonatomic, readwrite) int64_t onEvent;

@property(nonatomic, readwrite) BOOL ignoresScroll;

@end

#pragma mark - MatchaPBTouchButtonEvent

typedef GPB_ENUM(MatchaPBTouchButtonEvent_FieldNumber) {
  MatchaPBTouchButtonEvent_FieldNumber_Timestamp = 1,
  MatchaPBTouchButtonEvent_FieldNumber_Inside = 3,
  MatchaPBTouchButtonEvent_FieldNumber_Kind = 4,
};

@interface MatchaPBTouchButtonEvent : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) GPBTimestamp *timestamp;
/** Test to see if @c timestamp has been set. */
@property(nonatomic, readwrite) BOOL hasTimestamp;

@property(nonatomic, readwrite) BOOL inside;

@property(nonatomic, readwrite) MatchaPBTouchEventKind kind;

@end

/**
 * Fetches the raw value of a @c MatchaPBTouchButtonEvent's @c kind property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t MatchaPBTouchButtonEvent_Kind_RawValue(MatchaPBTouchButtonEvent *message);
/**
 * Sets the raw value of an @c MatchaPBTouchButtonEvent's @c kind property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetMatchaPBTouchButtonEvent_Kind_RawValue(MatchaPBTouchButtonEvent *message, int32_t value);

#pragma mark - MatchaPBTouchTapRecognizer

typedef GPB_ENUM(MatchaPBTouchTapRecognizer_FieldNumber) {
  MatchaPBTouchTapRecognizer_FieldNumber_Count = 1,
  MatchaPBTouchTapRecognizer_FieldNumber_RecognizedFunc = 2,
};

@interface MatchaPBTouchTapRecognizer : GPBMessage

@property(nonatomic, readwrite) int64_t count;

@property(nonatomic, readwrite) int64_t recognizedFunc;

@end

#pragma mark - MatchaPBTouchTapEvent

typedef GPB_ENUM(MatchaPBTouchTapEvent_FieldNumber) {
  MatchaPBTouchTapEvent_FieldNumber_Timestamp = 1,
  MatchaPBTouchTapEvent_FieldNumber_Position = 2,
  MatchaPBTouchTapEvent_FieldNumber_Kind = 3,
};

@interface MatchaPBTouchTapEvent : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) GPBTimestamp *timestamp;
/** Test to see if @c timestamp has been set. */
@property(nonatomic, readwrite) BOOL hasTimestamp;

@property(nonatomic, readwrite, strong, null_resettable) MatchaLayoutPBPoint *position;
/** Test to see if @c position has been set. */
@property(nonatomic, readwrite) BOOL hasPosition;

@property(nonatomic, readwrite) MatchaPBTouchEventKind kind;

@end

/**
 * Fetches the raw value of a @c MatchaPBTouchTapEvent's @c kind property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t MatchaPBTouchTapEvent_Kind_RawValue(MatchaPBTouchTapEvent *message);
/**
 * Sets the raw value of an @c MatchaPBTouchTapEvent's @c kind property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetMatchaPBTouchTapEvent_Kind_RawValue(MatchaPBTouchTapEvent *message, int32_t value);

#pragma mark - MatchaPBTouchPressRecognizer

typedef GPB_ENUM(MatchaPBTouchPressRecognizer_FieldNumber) {
  MatchaPBTouchPressRecognizer_FieldNumber_MinDuration = 1,
  MatchaPBTouchPressRecognizer_FieldNumber_FuncId = 2,
};

@interface MatchaPBTouchPressRecognizer : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) GPBDuration *minDuration;
/** Test to see if @c minDuration has been set. */
@property(nonatomic, readwrite) BOOL hasMinDuration;

@property(nonatomic, readwrite) int64_t funcId;

@end

#pragma mark - MatchaPBTouchPressEvent

typedef GPB_ENUM(MatchaPBTouchPressEvent_FieldNumber) {
  MatchaPBTouchPressEvent_FieldNumber_Timestamp = 1,
  MatchaPBTouchPressEvent_FieldNumber_Position = 2,
  MatchaPBTouchPressEvent_FieldNumber_Kind = 3,
  MatchaPBTouchPressEvent_FieldNumber_Duration = 4,
};

@interface MatchaPBTouchPressEvent : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) GPBTimestamp *timestamp;
/** Test to see if @c timestamp has been set. */
@property(nonatomic, readwrite) BOOL hasTimestamp;

@property(nonatomic, readwrite, strong, null_resettable) MatchaLayoutPBPoint *position;
/** Test to see if @c position has been set. */
@property(nonatomic, readwrite) BOOL hasPosition;

@property(nonatomic, readwrite) MatchaPBTouchEventKind kind;

@property(nonatomic, readwrite, strong, null_resettable) GPBDuration *duration;
/** Test to see if @c duration has been set. */
@property(nonatomic, readwrite) BOOL hasDuration;

@end

/**
 * Fetches the raw value of a @c MatchaPBTouchPressEvent's @c kind property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t MatchaPBTouchPressEvent_Kind_RawValue(MatchaPBTouchPressEvent *message);
/**
 * Sets the raw value of an @c MatchaPBTouchPressEvent's @c kind property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetMatchaPBTouchPressEvent_Kind_RawValue(MatchaPBTouchPressEvent *message, int32_t value);

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
