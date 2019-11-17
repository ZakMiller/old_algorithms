//
//  MagicNumberGenerator.cpp
//  12
//
//  Created by Zachary Miller on 4/10/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#include "TriangleNumberGenerator.hpp"

long TriangleNumberGenerator::getTriangleNumber(int n) const
{
        long total = 0;
        for(int i = 1; i <= n; i++)
        {
            total+= i;
        }
        
        return total;
}