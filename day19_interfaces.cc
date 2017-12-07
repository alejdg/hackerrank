//Write your code here
class Calculator : public AdvancedArithmetic {
    public:
        int divisorSum(int n) {
            int i = n;
            int sum = 0;
                
            while ( i > 1) {
                if ( n % i == 0 ) {
                    sum += i;
                }
                i--;
            }
            return sum+1;
        }
};
