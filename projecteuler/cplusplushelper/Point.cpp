//
//  Point.cpp
//  11
//
//  Created by Zachary Miller on 4/6/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#include "Point.hpp"


Point::Point(int x, int y)
{
    this->xCoord = x;
    this->yCoord = y;
}

int Point::x () const
{
    return this->xCoord;
}

int Point::y() const
{
    return this->yCoord;
}

bool Point::operator==(const Point& other) const
{
    if(x() == other.x() && y() == other.y())
        return true;
    else
        return false;
}

Point::Point(const Point &point)
{
    xCoord = point.x();
    yCoord = point.y();
}





