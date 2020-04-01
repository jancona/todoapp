part of todoclient.api;

class ModelToDoInput {
  
  bool completed = null;
  
  int order = null;
  
  String title = null;
  ModelToDoInput();

  @override
  String toString() {
    return 'ModelToDoInput[completed=$completed, order=$order, title=$title, ]';
  }

  ModelToDoInput.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    completed = json['completed'];
    order = json['order'];
    title = json['title'];
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (completed != null)
      json['completed'] = completed;
    if (order != null)
      json['order'] = order;
    if (title != null)
      json['title'] = title;
    return json;
  }

  static List<ModelToDoInput> listFromJson(List<dynamic> json) {
    return json == null ? List<ModelToDoInput>() : json.map((value) => ModelToDoInput.fromJson(value)).toList();
  }

  static Map<String, ModelToDoInput> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, ModelToDoInput>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = ModelToDoInput.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of ModelToDoInput-objects as value to a dart map
  static Map<String, List<ModelToDoInput>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<ModelToDoInput>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = ModelToDoInput.listFromJson(value);
       });
     }
     return map;
  }
}

