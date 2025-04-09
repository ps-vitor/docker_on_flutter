import 'package:http/http.dart' as http;
import 'dart:convert';

class ScrapService {
  final String ip;
  ScrapService(this.ip);

  Future<String> scrapUrl(String urlToScrap) async {
    final uri = Uri.parse("http://$ip:8080/scrape?url=$urlToScrap");

    try {
      final response = await http.get(uri);

      if (response.statusCode == 200) {
        return json.decode(response.body);
      } else {
        return 'Erro: ${response.statusCode}';
      }
    } catch (e) {
      return 'Error: $e';
    }
  }
}
