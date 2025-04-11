import 'package:http/http.dart' as http;
import 'dart:convert';

class ScrapService {
  final String ip;
  ScrapService(this.ip);

  Future<String> scrapUrl(String urlToScrap, scrapInterval) async {
    final uri = Uri.parse(
        "http://$ip:8080/add-job?url=$urlToScrap&interval=$scrapInterval");
    final scrapUri = Uri.parse("http://$ip:8080/srap?url=$urlToScrap");

    if (urlToScrap.isEmpty || scrapInterval.isEmpty) {
      return "URL or Interval must be filled";
    }

    try {
      final jobResponse = await http.get(uri);
      final scrapResponse = await http.get(scrapUri);
      final message = jsonDecode(scrapResponse.body)['message'];

      if (jobResponse.statusCode >= 200 && jobResponse.statusCode <= 201) {
        if (scrapResponse.statusCode == 200 ||
            scrapResponse.statusCode == 201) {
          return "Job add success";
        } else {
          return "Erro: $message";
        }
      } else {
        return "Erro: $message";
      }
    } catch (e) {
      return 'Error: $e';
    }
  }
}
