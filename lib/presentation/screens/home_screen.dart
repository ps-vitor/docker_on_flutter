import 'package:flutter/material.dart';
import '../../core/services/scraper_service.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text("Flutter + Rust Scraper")),
      
    );
  }
}
