import 'package:flutter/material.dart';
import 'package:file_picker/file_picker.dart';

import '../settings/settings_view.dart';
import 'package:dio/dio.dart';
// import 'dio';

class ReportPage extends StatelessWidget {
  static const routeName = '/';

  final dio = Dio();

  ReportPage({super.key});

  void getHttp() async {
    final response = await dio.get('http://127.0.0.1:4000/order?id=id_1');
    print(response);
  }

  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    return Scaffold(
      appBar: AppBar(
        actions: [
          IconButton(
            icon: const Icon(Icons.settings),
            onPressed: () {
              // Navigate to the settings page. If the user leaves and returns
              // to the app after it has been killed while running in the
              // background, the navigation stack is restored.
              Navigator.restorablePushNamed(context, SettingsView.routeName);
            },
          ),
        ],
      ),
      body: SafeArea(
          child: Padding(
        padding: const EdgeInsetsDirectional.fromSTEB(16, 12, 16, 0),
        child: SingleChildScrollView(
          child: Column(children: [
            const Text('Create Report'),
            const Padding(
              padding: EdgeInsetsDirectional.fromSTEB(0, 4, 0, 0),
            ),
            Column(
              children: [
                TextField(
                    decoration: InputDecoration(
                  labelText: 'Input name...',
                  enabledBorder: UnderlineInputBorder(
                    borderSide: const BorderSide(
                      color: Color(0xFFE5E7EB),
                      width: 2,
                    ),
                    borderRadius: BorderRadius.circular(0),
                  ),
                )),
                TextField(
                  autofocus: true,
                  obscureText: false,
                  decoration: InputDecoration(
                    hintText: 'Input report',
                    enabledBorder: UnderlineInputBorder(
                      borderSide: const BorderSide(
                        color: Color(0xFFE5E7EB),
                        width: 2,
                      ),
                      borderRadius: BorderRadius.circular(0),
                    ),
                    focusedBorder: UnderlineInputBorder(
                      borderSide: const BorderSide(
                        color: Color(0xFF6F61EF),
                        width: 2,
                      ),
                      borderRadius: BorderRadius.circular(0),
                    ),
                    errorBorder: UnderlineInputBorder(
                      borderSide: const BorderSide(
                        color: Color(0xFFFF5963),
                        width: 2,
                      ),
                      borderRadius: BorderRadius.circular(0),
                    ),
                    focusedErrorBorder: UnderlineInputBorder(
                      borderSide: const BorderSide(
                        color: Color(0xFFFF5963),
                        width: 2,
                      ),
                      borderRadius: BorderRadius.circular(0),
                    ),
                    contentPadding:
                        const EdgeInsetsDirectional.fromSTEB(16, 24, 16, 12),
                  ),
                  maxLines: 100,
                  minLines: 6,
                  cursorColor: const Color(0xFF6F61EF),
                )
              ],
            ),
            const Padding(
              padding: EdgeInsetsDirectional.fromSTEB(16, 12, 16, 0),
            ),
            Row(children: [
              IconButton(
                onPressed: () async {
                  FilePickerResult? result =
                      await FilePicker.platform.pickFiles();

                  if (result != null) {
                    print(result.files.first.name);
                  }
                },
                icon: const Icon(Icons.upload_file),
              ),
              IconButton(
                onPressed: () {
                  print("post");
                  getHttp();
                },
                icon: const Icon(Icons.outbond_outlined),
              )
            ]),
          ]),
        ),
      )),
    );
  }
}

class SentMailList extends StatelessWidget {
  final List<String> sentMails = [
    'Письмо 1',
    'Письмо 2',
    'Письмо 3',
    'Письмо 4',
    'Письмо 5',
  ];

 SentMailList({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Отправленные в почте'),
      ),
      drawer: Drawer(
        child: Column(
          children: <Widget>[
            Container(
              width: double.infinity, // Растянуть на всю ширину
              child: DrawerHeader(
                decoration: BoxDecoration(
                  color: Colors.blue,
                ),
                child: Text(
                  'Меню',
                  style: TextStyle(
                    color: Colors.white,
                    fontSize: 24,
                  ),
                ),
              ),
            ),
            ListTile(
              title: const Text('Отправленные'),
              onTap: () {
                Navigator.pop(context); // Закрыть боковую панель
              },
            ),
            ListTile(
              title: const Text('Ожидают отправки'),
              onTap: () {
                Navigator.pop(context); // Закрыть боковую панель
                Navigator.push(
                  context,
                  MaterialPageRoute(
                    builder: (context) => PendingMailList(),
                  ),
                );
              },
            ),
            Expanded(
              child: Align(
                alignment: Alignment.bottomCenter,
                child: Container(
                  color: Colors.orange, // Ваш цвет
                  child: ListTile(
                    title: const Text(
                      'Создать сообщение',
                      style: TextStyle(
                        color: Colors.white,
                      ),
                    ),
                    onTap: () {
                      Navigator.pop(context); // Закрыть боковую панель
                      Navigator.push(
                        context,
                        MaterialPageRoute(
                          builder: (context) => ReportPage(),
                        ),
                      );
                    },
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
      body: ListView.builder(
        itemCount: sentMails.length,
        itemBuilder: (context, index) {
          return ListTile(
            title: Text(sentMails[index]),
            onTap: () {
              // Действие при нажатии на элемент списка (по желанию)
              // Например, можно открыть детали письма или выполнить другое действие
              print('Выбрано письмо: ${sentMails[index]}');
            },
          );
        },
      ),
    );
  }
}

class PendingMailList extends StatelessWidget {
  final List<String> pendingMails = [
    'Письмо A',
    'Письмо B',
    'Письмо C',
    'Письмо D',
    'Письмо E',
  ];

PendingMailList({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Ожидают отправки'),
      ),
      body: ListView.builder(
        itemCount: pendingMails.length,
        itemBuilder: (context, index) {
          return ListTile(
            title: Text(pendingMails[index]),
            onTap: () {
              // Действие при нажатии на элемент списка (по желанию)
              // Например, можно открыть детали письма или выполнить другое действие
              print('Выбрано письмо: ${pendingMails[index]}');
            },
          );
        },
      ),
    );
  }
}
