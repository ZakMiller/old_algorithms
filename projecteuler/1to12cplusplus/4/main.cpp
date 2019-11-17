//
//  main.cpp
//  4
//
//  Created by Zachary Miller on 4/2/16.
//  Copyright © 2016 Spllongerfloss. All rights reserved.
//

#include <iostream>
class PalindromeTester{
public:
    bool isPalindrome(long n)
    {
        long numDigitsToCompare = numDigits(n) / 2;
        for(long i = 0; i<numDigitsToCompare; i++)
            if(!areDigitsTheSameForNumberAtOffset(n, i))
                return false;
        return true;
    }
private:
    long numDigits(long n)
    {
        long count = 0;
        while(n > 0)
        {
            n /= 10;
            count++;
        }
        return count;
    }
    
    bool areDigitsTheSameForNumberAtOffset(long n, long offset)
    {
        return getNumberAtIndex(n, offset) == getNumberAtIndex(n, numDigits(n)-offset-1);
            
    }
    
    long getNumberAtIndex(long n, long index)
    {
        for(long i = 0; i<index; i++)
            n /= 10;
        
        return n%10;
    }
    

    
};


int main(int argc, const char * argv[]) {
    PalindromeTester palindromeTester;
    long highestPalindrome = 0;
    for(int i = 999; i>99; --i)
        for(int j = 999; j>=i; --j)
    {
            long possiblePalindrome = i*j;
            if(possiblePalindrome < highestPalindrome)
                continue;
            if(palindromeTester.isPalindrome(possiblePalindrome))
            {
                highestPalindrome = std::max(highestPalindrome, possiblePalindrome);
                
            }
    }
    std::cout << highestPalindrome;
    return 0;
    

}


