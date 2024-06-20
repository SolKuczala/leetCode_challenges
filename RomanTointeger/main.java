// ""MCMXCIV""
class Solution {
    public int romanToInt(String s) {
        HashMap<String, Integer> Numbers = new HashMap<String, Integer>();

        Numbers.put("I", 1);
        Numbers.put("V", 5);
        Numbers.put("X", 10);
        Numbers.put("L", 50);
        Numbers.put("C", 100);
        Numbers.put("D", 500);
        Numbers.put("M", 1000);

        String numberRoman = s;

        int rezult = Numbers.get(Character.toString(numberRoman.charAt(numberRoman.length() -1)));
        for (int i = numberRoman.length() - 2; i >= 0;i-- ){
            if (Numbers.get(Character.toString(numberRoman.charAt(i))) < Numbers.get(Character.toString(numberRoman.charAt(i+1))) ){
                rezult -= Numbers.get(Character.toString(numberRoman.charAt(i)))
            } else {
                rezult += Numbers.get(Character.toString(numberRoman.charAt(i)))
            }
        }
        return rezult;
    }
}

/*
if (i != 0) {
                if (Numbers.get(Character.toString(numberRoman.charAt(i))) <= Numbers.get(Character.toString(numberRoman.charAt(i-1)))){
                    cont += Numbers.get(Character.toString(numberRoman.charAt(i)));
                }else{
                    if (i == 1){
                        cont += Numbers.get(Character.toString(numberRoman.charAt(i))) - Numbers.get(Character.toString(numberRoman.charAt(i-1)));
                        break;
                    }
                    cont += Numbers.get(Character.toString(numberRoman.charAt(i))) - Numbers.get(Character.toString(numberRoman.charAt(i-1)));
                    i--;
                }
            }
            if (i == 0) {
                cont += Numbers.get(Character.toString(numberRoman.charAt(i)));
            }
*/ */