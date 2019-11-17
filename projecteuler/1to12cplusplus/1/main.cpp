//
//  main.cpp
//  1
//
//  Created by Zachary Miller on 4/1/16.
//  Copyright © 2016 Splinterfloss. All rights reserved.
//

#include <iostream>

int main(int argc, const char * argv[]) {
    int sum = 0;
    
    for(int i=0; i<1000; i++)
        if(!(i%3) || !(i%5))
            sum += i;
    std::cout << sum;
    return 0;
}
