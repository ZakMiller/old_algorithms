//
//  main.cpp
//  6
//
//  Created by Zachary Miller on 4/3/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#include <iostream>
#include <cmath>
class SquaresAndSums{
public:
    int findSquareOfSumsMinusSumOfSquaresOneTo(int n)
    {
        return findSquareOfSumsOneTo(n) - findSumOfSquaresOneTo(n);
    }
private:
    int findSquareOfSumsOneTo(int n){
        int total = 0;
        for(int i=1; i<=n; i++)
            total+= i;
        return std::pow(total, 2);
        
    }
    int findSumOfSquaresOneTo(int n){
        int total = 0;
        for(int i=1; i<=n; i++)
            total += std::pow(i, 2);
        return total;
            
        
    }
};
int main(int argc, const char * argv[]) {

    SquaresAndSums squaresAndSums;
    std:: cout << squaresAndSums.findSquareOfSumsMinusSumOfSquaresOneTo(100);
    return 0;
}
