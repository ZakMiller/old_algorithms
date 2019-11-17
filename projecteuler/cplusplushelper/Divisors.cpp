//
//  Divisors.cpp
//  12
//
//  Created by Zachary Miller on 4/10/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#include "Divisors.hpp"
long Divisors::getNumDivisors(long n) const
{
    long numDivisors = 2;
    for (int i = 2; i <= n/2; i++)
    {
        if( n % i == 0)
        {
            numDivisors++;
        }
    }
    return numDivisors;
}