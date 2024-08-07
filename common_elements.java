/*
 * Running:
 *      # Build
 *      /usr/local/Cellar/openjdk/21.0.2/bin/javac common_elements.java
 *      # Run
 *      /usr/local/Cellar/openjdk/21.0.2/bin/java common_elements.java 
 */

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;

class PersonRecord {
    String id;
    String name;
    public PersonRecord(String id, String name){
        this.id = id;
        this.name = name;
    }
}

class CsvReader {
    public ArrayList<HashMap<String, String>> reader(String csvFile){
        String line = "";
        String cvsSplitBy = ",";
        ArrayList list = new ArrayList<HashMap<String, String>>();
        try (BufferedReader br = new BufferedReader(new FileReader(csvFile))) {

            String headers[] = null;

            while ((line = br.readLine()) != null) {

                if(headers == null){
                    headers = line.split(",");
                } else {
                    String[] elements = line.split(",");

                    HashMap<String, String> hm = new HashMap<String, String>();
                    for(int i = 0; i < Math.min(elements.length, headers.length); i++){
                        hm.put(headers[i], elements[i]);
                    }
                    list.add(hm);
                }
            }


        } catch (IOException e) {
            e.printStackTrace();
        }

        return list;
    }
}

class CommonElements {

    public ArrayList<String> commonElements(ArrayList<HashMap<String, String>> left, ArrayList<HashMap<String, String>> right){
        ArrayList list = new ArrayList<String>();

        for(HashMap<String, String> l : left){
            for(HashMap<String, String> r : right){
                if(l.equals(r)){
                    list.add(l.get("id"));
                } else {
                }
            }
        }
        return list;    
    }

}

public class common_elements {

    public static void main(String[] args) {

        CsvReader csvReader = new CsvReader();
        ArrayList left = csvReader.reader("left.csv");
        ArrayList right = csvReader.reader("right.csv");
        System.out.println(left);
        System.out.println(right);


        CommonElements commonElements = new CommonElements();
        ArrayList<String> output = commonElements.commonElements(left, right);
        System.out.println(output);

    }

}
