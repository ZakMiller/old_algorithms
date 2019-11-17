//
//  main.cpp
//  3
//
//  Created by Zachary Miller on 4/2/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#include <iostream>
class LargestPrimeFactor
{
    public: long getLargestPrimeFactorOf(long n)
    {
        for(long i=2; i<n/2;i++){
            if(n%i!=0)
                continue;
            long factor = n / i;
            if(isPrime(factor))
                return factor;
        }
        return -1;
    }
    
    bool isPrime(long n)
    {
        if(n%2==0)
            return false;
        
        for(long i=3; i<n/2; i+=2)
            if(n%i==0)
                return false;
        
        return true;
    }
};

int main(int argc, const char * argv[]) {
    LargestPrimeFactor largestPrimeFactor;
    std::cout << largestPrimeFactor.getLargestPrimeFactorOf(600851475143);
    return 0;
}


