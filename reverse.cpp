#include <iostream>
#include <climits>

using namespace std;

class Solution {
public:
    int reverse(int x) {
        int reversed = 0;
        while (x != 0) {
            int last = x % 10;

            if (reversed > INT_MAX / 10 || (reversed == INT_MAX / 10 && last > 7)) {
                return 0; 
            }
            if (reversed < INT_MIN / 10 || (reversed == INT_MIN / 10 && last < -8)) {
                return 0; 
            }

            reversed = reversed * 10 + last;
            x /= 10;
        }
        return reversed;
    }
};


int main(){
Solution solution;
cout<<solution.reverse(123);
}