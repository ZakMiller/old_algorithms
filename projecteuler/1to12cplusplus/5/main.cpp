//
//  main.cpp
//  5
//
//  Created by Zachary Miller on 4/3/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#include <iostream>
class EvenlyDivisibleChecker{
public: bool isNumberEvenlyDisibleByOneTo(int number, int divisorCap)
    {
        for(int i=2; i<=divisorCap; i++)
            if(number % i != 0)
                return false;
        return true;
    }
    
};

int main(int argc, const char * argv[]) {
    EvenlyDivisibleChecker evenlyDivisibleChecker;
    int numberToCheck = 20;
    while(true)
    {
        if(evenlyDivisibleChecker.isNumberEvenlyDisibleByOneTo(numberToCheck, 20)){
            std::cout << numberToCheck;
            return 0;
        }
        ++numberToCheck;
    }
        
}


