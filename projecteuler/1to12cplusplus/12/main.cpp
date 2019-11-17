//
//  main.cpp
//  12
//
//  Created by Zachary Miller on 4/10/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#include <iostream>
#include "TriangleNumberGenerator.hpp"
#include "Divisors.hpp"

int main(int argc, const char * argv[]) {
    TriangleNumberGenerator triangleNumberGenerator;
    Divisors divisors;
    long numDivisors = 0;
    long triangleNumber = 0;
    int n = 1;
    while (numDivisors <= 500)
    {
        triangleNumber = triangleNumberGenerator.getTriangleNumber(n);
        numDivisors = divisors.getNumDivisors(triangleNumber);
        n++;
    }
    std::cout << triangleNumber;
    return 0;
}