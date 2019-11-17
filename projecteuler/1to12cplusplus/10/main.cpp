//
//  main.cpp
//  10
//
//  Created by Zachary Miller on 4/5/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#include <iostream>
class PrimeFinder{
public:
    long sumPrimesUpToValue(int n)
    {
        long total= 0;
        for(int i=2; i<n; i++)
            if(isNumberPrime(i))
                total += i;
        return total;
                
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
    std::cout << primeFinder.sumPrimesUpToValue(2000000);
    return 0;
}
