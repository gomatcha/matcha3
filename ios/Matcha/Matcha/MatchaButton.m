#import "MatchaButton.h"
#import "MatchaViewController.h"
#import "MatchaProtobuf.h"

@interface MatchaButton ()
@property (nonatomic, strong) UIButton *button;
@property (nonatomic, weak) MatchaViewNode *viewNode;
@property (nonatomic, strong) MatchaNode *node;
@end

@implementation MatchaButton

+ (void)load {
    MatchaRegisterView(@"gomatcha.io/matcha/view/button", ^(MatchaViewNode *node){
        return [[MatchaButton alloc] initWithViewNode:node];
    });
}

- (id)initWithViewNode:(MatchaViewNode *)viewNode {
    if ((self = [super initWithFrame:CGRectZero])) {
        self.viewNode = viewNode;
        self.button = [UIButton buttonWithType:UIButtonTypeSystem];
        [self.button addTarget:self action:@selector(onPress) forControlEvents:UIControlEventTouchUpInside];
        [self addSubview:self.button];
    }
    return self;
}

- (void)setNode:(MatchaNode *)value {
    _node = value;
    GPBAny *state = value.nativeViewState;
    NSError *error = nil;
    MatchaButtonPBView *pbbutton = (id)[state unpackMessageClass:[MatchaButtonPBView class] error:&error];
    
    NSAttributedString *string = [[NSAttributedString alloc] initWithProtobuf:pbbutton.styledText];
    [self.button setAttributedTitle:string forState:UIControlStateNormal];
}

- (void)layoutSubviews {
    self.button.frame = self.bounds;
}

- (void)onPress {
    [self.viewNode.rootVC call:@"OnPress" viewId:self.node.identifier.longLongValue args:@[]];
}

@end
