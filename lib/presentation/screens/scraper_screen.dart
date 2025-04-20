import 'package:flutter/material.dart';
import 'package:docker_on_flutter/core/services/scraper_service.dart';

class ScraperScreen extends StatefulWidget {
  @override
  _ScraperScreenState createState() => _ScraperScreenState();
}

class _ScraperScreenState extends State<ScraperScreen> {
  final TextEditingController _controller = TextEditingController();
  final intervalController = TextEditingController();
  String result = "";

  void fetchData() async {
    final url = _controller.text.trim();
    final interval = intervalController.text.trim();
    if (url.isEmpty) return;

    final service = ScrapService("192.168.1.14");
    final response = await service.scrapUrl(url, interval);

    setState(() {
      result = response;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Web Scraper")),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(children: [
          TextField(
            controller: _controller,
            decoration: InputDecoration(
              labelText: 'URL to scrap',
              border: OutlineInputBorder(),
            ),
          ),
          TextField(
            controller: intervalController,
            decoration: InputDecoration(
              labelText: 'Interval (seconds)',
              border: OutlineInputBorder(),
            ),
          ),
          SizedBox(
            height: 12,
          ),
          ElevatedButton(onPressed: fetchData, child: Text("Scrap")),
          SizedBox(height: 20),
          Expanded(
              child: SingleChildScrollView(
            child: Text(result),
          ))
        ]),
      ),
    );
  }
}
