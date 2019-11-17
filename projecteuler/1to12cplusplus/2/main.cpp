//
//  main.cpp
//  2
//
//  Created by Zachary Miller on 4/2/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#include <iostream>
class FibCompute
{
    int sum = 0;

public: int getSumForFibUpTo(int cap)
    {
        int previousFib = 1;
        int currentFib = 1;
        int temp = 0;
        while(currentFib < cap)
        {
            if(currentFib % 2 == 0)
                sum+= currentFib;
            
            temp = currentFib;
            currentFib += previousFib;
            previousFib = temp;
            
        }
        return sum;
    }
};


int main(int argc, const char * argv[]) {
    FibCompute fibCompute;
    std::cout << fibCompute.getSumForFibUpTo(5000000);
    
}


