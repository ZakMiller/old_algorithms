//
//  Grid.hpp
//  11
//
//  Created by Zachary Miller on 4/6/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#ifndef Grid_hpp
#define Grid_hpp

#include <stdio.h>
#include <vector>
#include <map>
#include <set>

#include "Point.hpp"
class Grid
{
    std::vector<std::vector<int>> arr;
public: enum Direction { UP, UPRIGHT, RIGHT, DOWNRIGHT, DOWN, DOWNLEFT, LEFT, UPLEFT, DirectionFirst = UP, DirectionLast = UPLEFT };
public:
    Grid(std::vector<std::vector<int>>);
    int getValueAtPoint(const Point &p) const;
    Point getAdjacentPoint(int x, int y, Direction direction) const;
    Point getAdjacentPoint(const Point &p, Direction direction) const;
    std::vector<const Point> getValidAdjacentPoints(int x, int y) const;

    size_t length() const;
    size_t height() const;
};
#endif /* Grid_hpp */