#import <UIKit/UIKit.h>
#import <MatchaBridge/MatchaBridge.h>

@interface MatchaObjcBridge (Extensions)
- (void)configure;
- (MatchaGoValue *)sizeForAttributedString:(NSData *)data;
- (void)updateId:(NSInteger)identifier withProtobuf:(NSData *)protobuf;
- (NSString *)assetsDir;
- (MatchaGoValue *)imageForResource:(NSString *)path;
- (MatchaGoValue *)propertiesForResource:(NSString *)path;
@end
