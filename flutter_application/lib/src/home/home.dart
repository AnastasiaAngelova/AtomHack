import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';

import '../settings/settings_view.dart';
import 'package:dio/dio.dart';
// import 'dio';

class ReportPage extends StatelessWidget {
  static const routeName = '/';

  final dio = Dio();

  void getHttp() async {
    try {
      final response = await dio.get('http://127.0.0.1:4000/order?id=id_1');
      print(response);
    } catch (e) {
      print(e);
    }
  }

  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    return Scaffold(
        appBar: AppBar(
          title:
              Text("MARS report system", style: TextStyle(color: Colors.white)),
          backgroundColor: Colors.deepPurple[900],
          actions: [
            IconButton(
              icon: const Icon(
                  color: Color.fromARGB(255, 255, 255, 255), Icons.settings),
              onPressed: () {
                // Navigate to the settings page. If the user leaves and returns
                // to the app after it has been killed while running in the
                // background, the navigation stack is restored.
                Navigator.restorablePushNamed(context, SettingsView.routeName);
              },
            ),
            IconButton(
              icon: const Icon(
                  color: Color.fromARGB(255, 255, 255, 255),
                  Icons.replay_outlined),
              onPressed: () {
                // Navigate to the settings page. If the user leaves and returns
                // to the app after it has been killed while running in the
                // background, the navigation stack is restored.
                this.build(context);
              },
            ),
          ],
        ),
        body: Row(children: [Left(), Middle()]));
  }
}

class Middle extends StatelessWidget {
  const Middle({super.key});
  // final List<Report> items;

  @override
  Widget build(BuildContext context) {
    return ColoredBox(
      color: Colors.black12,
      // child: ListView.builder(
      //     restorationId: 'sampleItemListView',
      //     // itemCount: 3,
      //     itemBuilder: (BuildContext context, int index) {}),
    );
  }
}

class Report extends StatelessWidget {
  const Report({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Container();
  }
}

class Left extends StatelessWidget {
  const Left({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return ColoredBox(
      color: Colors.deepPurple,
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Stack(
          children: [
            Column(
              children: [
                TextButton.icon(
                  onPressed: () {
                    print("Отправлены");
                  },
                  icon: Icon(
                    color: Color.fromARGB(255, 255, 255, 255),
                    Icons.send_rounded,
                  ),
                  label:
                      Text("Отправлены", style: TextStyle(color: Colors.white)),
                ),
                TextButton.icon(
                    onPressed: () {
                      print("Ожидают отправки");
                    },
                    icon: Icon(
                      Icons.timelapse,
                      color: Color.fromARGB(255, 255, 255, 255),
                    ),
                    label: Text("Ожидают отправки",
                        style: TextStyle(color: Colors.white))),
                TextButton.icon(
                    onPressed: () {
                      print("Edit draft");
                    },
                    icon: Icon(
                      Icons.edit_document,
                      color: Color.fromARGB(255, 255, 255, 255),
                    ),
                    label: Text("Черновик",
                        style: TextStyle(color: Colors.white))),
              ],
            ),
            Padding(
                padding: const EdgeInsets.symmetric(vertical: Checkbox.width)),
            Column(
              mainAxisAlignment: MainAxisAlignment.end,
              children: [
                TextButton.icon(
                    onPressed: () {
                      print("Create report");
                    },
                    icon: Icon(
                      Icons.note_add,
                      color: Color.fromARGB(255, 255, 255, 255),
                    ),
                    label: Text("Create report",
                        style: TextStyle(color: Colors.white))),
              ],
            ),
          ],
        ),
      ),
    );
  }
}

/*
TextButton.icon(
              onPressed: () {
                print("pressed1");
              },
              icon: Icon(Icons.upload_file),
              label: Text("Upload File")),
          TextButton.icon(
              onPressed: () {
                print("pressed2");
              },
              icon: Icon(Icons.upload_file),
              label: Text("Push report"))
*/
