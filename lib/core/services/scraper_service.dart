import 'package:http/http.dart' as http;
import 'dart:convert';

class ScrapService {
  final String ip;
  ScrapService(this.ip);

  Future<String> scrapUrl(String urlToScrap, scrapInterval) async {
    if (urlToScrap.isEmpty || scrapInterval.isEmpty) {
      return "URL or Interval must be filled";
    }

    final uri = Uri.parse(
        "http://$ip:8080/add-job?url=$urlToScrap&interval=$scrapInterval");
    final scrapUri = Uri.parse("http://$ip:8080/scrape?url=$urlToScrap");

    try {
      final jobResponse = await http.get(uri);
      final scrapResponse = await http.get(scrapUri);

      if (jobResponse.statusCode == 200 && scrapResponse.statusCode == 200) {
        final message = jsonDecode(scrapResponse.body)['message'];
        return "Job add success: $message";
      } else {
        return "Erro: ${jobResponse.statusCode} / ${scrapResponse.statusCode}\n${scrapResponse.body}";
      }
    } catch (e) {
      return 'Error: $e';
    }
  }
}
