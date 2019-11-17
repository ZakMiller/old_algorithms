//
//  Point.hpp
//  11
//
//  Created by Zachary Miller on 4/6/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#ifndef Point_hpp
#define Point_hpp

#include <stdio.h>
class Point {
int xCoord;
int yCoord;
public:
    Point(int x, int  y);
    Point(const Point &point);
    int x() const;
    int y() const;
    bool operator==(const Point& other) const;
    
    
};
#endif /* Point_hpp */
