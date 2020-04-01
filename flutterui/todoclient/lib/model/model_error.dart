part of todoclient.api;

class ModelError {
  
  int code = null;
  
  String message = null;
  ModelError();

  @override
  String toString() {
    return 'ModelError[code=$code, message=$message, ]';
  }

  ModelError.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    code = json['code'];
    message = json['message'];
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (code != null)
      json['code'] = code;
    if (message != null)
      json['message'] = message;
    return json;
  }

  static List<ModelError> listFromJson(List<dynamic> json) {
    return json == null ? List<ModelError>() : json.map((value) => ModelError.fromJson(value)).toList();
  }

  static Map<String, ModelError> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, ModelError>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = ModelError.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of ModelError-objects as value to a dart map
  static Map<String, List<ModelError>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<ModelError>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = ModelError.listFromJson(value);
       });
     }
     return map;
  }
}

