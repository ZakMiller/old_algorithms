//
//  main.cpp
//  9
//
//  Created by Zachary Miller on 4/5/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#include <iostream>
#include <cmath>
class PythagoreanTriplet{
public:
    int findTripletWhoseSumIs(int n)
    {
        for(float a=100; a<500;a++)
        for(float b = a; b<600; b++)
        {
            float c = std::sqrt(std::pow(a, 2) + std::pow(b,2));
            if(ceilf(c) == c)
                if(a + b + c == n)
                    return a*b*c;
        }
        return -1;
    }
    
};
int main(int argc, const char * argv[]) {
    PythagoreanTriplet pythagoreanTriplet;
    std::cout << pythagoreanTriplet.findTripletWhoseSumIs(1000);
    return 0;
}
