package astros;

import java.net.HttpURLConnection;
import java.net.URL;
import java.util.Scanner;
import com.google.gson.Gson;

public class App {
    public static void main(String[] args) {
        String url = "http://api.open-notify.org/astros.json";
        People people = fetchAstros(url);
        System.out.format("there are %d astronauts currently in space%n", people.getNumber());
    }

    static People fetchAstros(String apiUrl) {
        try {
            // create a url from our input
            URL url = new URL(apiUrl);

            // create and execute an HTTP request
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("GET");
            conn.setRequestProperty("User-Agent", "spacecount-tutorial");
            conn.connect();

            // check the response code
            int responsecode = conn.getResponseCode();
            if (responsecode != 200) {
                throw new RuntimeException("HttpResponseCode: " + responsecode);
            } else {

                // Write all the JSON data into a string using a scanner
                String inline = "";
                Scanner scanner = new Scanner(url.openStream());
                while (scanner.hasNext()) {
                    inline += scanner.nextLine();
                }
                scanner.close();

                // unmarshal the JSON data
                Gson gson = new Gson();
                People people = gson.fromJson(inline, People.class);
                return people;
            }

            // check for errors
        } catch (Exception e) {
            e.printStackTrace();
        }
        return null;
    }
}
