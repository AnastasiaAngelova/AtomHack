import 'package:flutter/material.dart';

import '../settings/settings_view.dart';
import 'package:dio/dio.dart';
// import 'dio';

class ReportPage extends StatelessWidget {
  static const routeName = '/';

  final dio = Dio();

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
        padding: EdgeInsetsDirectional.fromSTEB(16, 12, 16, 0),
        child: SingleChildScrollView(
          child: Column(children: [
            Text('Create Report'),
            Padding(
              padding: EdgeInsetsDirectional.fromSTEB(0, 4, 0, 0),
            ),
            Column(
              children: [
                TextField(
                    decoration: InputDecoration(
                  labelText: 'Input name...',
                  enabledBorder: UnderlineInputBorder(
                    borderSide: BorderSide(
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
                      borderSide: BorderSide(
                        color: Color(0xFFE5E7EB),
                        width: 2,
                      ),
                      borderRadius: BorderRadius.circular(0),
                    ),
                    focusedBorder: UnderlineInputBorder(
                      borderSide: BorderSide(
                        color: Color(0xFF6F61EF),
                        width: 2,
                      ),
                      borderRadius: BorderRadius.circular(0),
                    ),
                    errorBorder: UnderlineInputBorder(
                      borderSide: BorderSide(
                        color: Color(0xFFFF5963),
                        width: 2,
                      ),
                      borderRadius: BorderRadius.circular(0),
                    ),
                    focusedErrorBorder: UnderlineInputBorder(
                      borderSide: BorderSide(
                        color: Color(0xFFFF5963),
                        width: 2,
                      ),
                      borderRadius: BorderRadius.circular(0),
                    ),
                    contentPadding:
                        EdgeInsetsDirectional.fromSTEB(16, 24, 16, 12),
                  ),
                  maxLines: 100,
                  minLines: 6,
                  cursorColor: Color(0xFF6F61EF),
                )
              ],
            ),
            Padding(
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

/**
Padding(
  padding: EdgeInsetsDirectional.fromSTEB(0, 24, 0, 12),
  child: FFButtonWidget(
    onPressed: () {
      print('Button pressed ...');
    },
    text: 'Submit Ticket',
    icon: Icon(
      Icons.receipt_long,
      size: 15,
    ),
    options: FFButtonOptions(
      width: double.infinity,
      height: 54,
      padding: EdgeInsets.all(0),
      iconPadding: EdgeInsetsDirectional.fromSTEB(0, 0, 0, 0),
      color: Color(0xFF6F61EF),
      textStyle: FlutterFlowTheme.of(context).titleSmall.override(
            fontFamily: 'Plus Jakarta Sans',
            color: Colors.white,
            fontSize: 16,
            fontWeight: FontWeight.w500,
          ),
      elevation: 4,
      borderSide: BorderSide(
        color: Colors.transparent,
        width: 1,
      ),
      borderRadius: BorderRadius.circular(12),
    ),
  ),
)
// Padding(
            //     padding: EdgeInsetsDirectional.fromSTEB(0, 16, 0, 0),
            //     child: Container(
            //       width: double.infinity,
            //       constraints: BoxConstraints(
            //         maxWidth: 500,
            //       ),
            //       decoration: BoxDecoration(
            //         // color: Colors.white,
            //         borderRadius: BorderRadius.circular(12),
            //         border: Border.all(
            //           // color: Color(0xFFE5E7EB),
            //           width: 2,
            //         ),
            //       ),
            //       child: Padding(
            //         padding: EdgeInsets.all(8),
            //         child: Row(
            //           mainAxisSize: MainAxisSize.max,
            //           children: [
            //             Icon(
            //               Icons.upload_file,
            //               color: Color(0xFF6F61EF),
            //               size: 32,
            //             ),
            //             Padding(
            //               padding: EdgeInsetsDirectional.fromSTEB(16, 0, 0, 0),
            //               child: Text(
            //                 'Upload File',
            //                 textAlign: TextAlign.center,
            //               ),
            //             ),
            //           ],
            //         ),
            //       ),
            //     )),
            // Padding(
            //     padding: EdgeInsetsDirectional.fromSTEB(0, 16, 0, 0),
            //     child: Container(
            //       width: double.infinity,
            //       constraints: BoxConstraints(
            //         maxWidth: 500,
            //       ),
            //       decoration: BoxDecoration(
            //         // color: Colors.white,
            //         borderRadius: BorderRadius.circular(12),
            //         border: Border.all(
            //           // color: Color(0xFFE5E7EB),
            //           width: 2,
            //         ),
            //       ),
            //       child: Padding(
            //         padding: EdgeInsets.all(8),
            //         child: Row(
            //           mainAxisSize: MainAxisSize.max,
            //           children: [
            //             Icon(
            //               Icons.airline_stops_sharp,
            //               color: Color(0xFF6F61EF),
            //               size: 32,
            //             ),
            //             Padding(
            //               padding: EdgeInsetsDirectional.fromSTEB(16, 0, 0, 0),
            //               child: Text(
            //                 'Submit',
            //                 textAlign: TextAlign.center,
            //               ),
            //             ),
            //           ],
            //         ),
            //       ),
            //     )),
            // Padding(
            //   padding: EdgeInsetsDirectional.fromSTEB(16, 12, 16, 0),
            // ),
 */
