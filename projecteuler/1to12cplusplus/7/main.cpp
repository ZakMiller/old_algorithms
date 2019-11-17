//
//  main.cpp
//  7
//
//  Created by Zachary Miller on 4/4/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#include <iostream>
class PrimeFinder
{
public:
    int getPrime(int n)
    {
        int possiblePrime = 2;
        int numberOfPrimes = 0;
        while(true)
        {
            if(isNumberPrime(possiblePrime)){
                numberOfPrimes++;
                if(numberOfPrimes==n)
                    return possiblePrime;
            }
            possiblePrime++;
            
        }
        return -1;
    }
private:
    bool isNumberPrime(int n)
    {
        for(int i=2; i<n; i++)
            if(n%i==0)
                return false;
        return true;
    }
};
int main(int argc, const char * argv[]) {
    PrimeFinder primeFinder;
    std::cout << primeFinder.getPrime(10001);
    return 0;
}
