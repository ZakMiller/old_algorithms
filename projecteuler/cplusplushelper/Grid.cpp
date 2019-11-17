//
//  Grid.cpp
//  11
//
//  Created by Zachary Miller on 4/6/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#include "Grid.hpp"
Grid::Grid(std::vector<std::vector<int>> arr)
{
    for(int i=0;i<arr.size();i++){
        this->arr.push_back(std::vector<int>());
        for(int j=0;j<arr[0].size();j++){
            this->arr[i].push_back(arr[i][j]);
        }
    }
}

int Grid::getValueAtPoint(const Point &p) const
{
    return this->arr[p.x()][p.y()];
}

Point Grid::getAdjacentPoint(int x, int y, Direction direction) const
{
    switch(direction)
    {
        case UP: return Point(x, y-1);
        case UPRIGHT: return Point(x+1, y-1);
        case RIGHT: return Point(x+1, y);
        case DOWNRIGHT: return Point(x+1, y+1);
        case DOWN: return Point(x, y+1);
        case DOWNLEFT: return Point(x-1, y+1);
        case LEFT: return Point(x-1, y);
        case UPLEFT: return Point(x-1, y-1);
    
    }
}

Point Grid::getAdjacentPoint(const Point &p, Direction direction) const
{
    return getAdjacentPoint(p.x(), p.y(), direction);
    }

std::vector<const Point> Grid::getValidAdjacentPoints(int x, int y) const
{
    std::vector<const Point> validAdjacentPoints;
    
    for(int direction=DirectionFirst; direction<= DirectionLast; direction++){
        const Point possiblePoint = getAdjacentPoint(x, y, static_cast<Direction>(direction));
        int x = possiblePoint.x();
        int y = possiblePoint.y();
        if(x > 0 && x < this->length() && y > 0 && y < this->height())
            validAdjacentPoints.push_back(possiblePoint);
    }
    return validAdjacentPoints;
}


size_t Grid::length() const
{
    return this->arr.size();
}

size_t Grid::height() const
{
    return this->arr[0].size();
}

