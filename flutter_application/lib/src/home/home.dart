import 'package:flutter/material.dart';

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
                onPressed: () {
                  print("pressed");
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
        child: ListView(
          padding: EdgeInsets.zero,
          children: <Widget>[
            const DrawerHeader(
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
                    builder: (context) =>  PendingMailList(),
                  ),
                );
              },
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
              Navigator.of(context).push(_buildPageRoute(sentMails[index]));
            },
          );
        },
      ),
    );
  }

  PageRouteBuilder _buildPageRoute(String emailContent) {
    return PageRouteBuilder(
      pageBuilder: (context, animation, secondaryAnimation) {
        return _EmailContentDialog(emailContent: emailContent);
      },
      transitionsBuilder: (context, animation, secondaryAnimation, child) {
        const begin = Offset(1.0, 0.0);
        const end = Offset.zero;
        const curve = Curves.easeInOut;
        var tween = Tween(begin: begin, end: end).chain(CurveTween(curve: curve));
        var offsetAnimation = animation.drive(tween);
        return SlideTransition(position: offsetAnimation, child: child);
      },
    );
  }
}

class _EmailContentDialog extends StatelessWidget {
  final String emailContent;

  const _EmailContentDialog({Key? key, required this.emailContent}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Email Content'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Text(emailContent),
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
