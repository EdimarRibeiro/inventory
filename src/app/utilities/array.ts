// http://www.dotnetspeak.com/typescript/extending-javascript-arrays-with-typescript/

declare global {
  interface Array<T> {
    firstOrDefault(predicate: (item: T) => boolean): T;
    where(predicate: (item: T) => boolean): T[];
    orderBy(propertyExpression: (item: T) => any): T[];
    orderByDescending(propertyExpression: (item: T) => any): T[];
    orderByMany(propertyExpressions: [(item: T) => any]): T[];
    orderByManyDescending(propertyExpressions: [(item: T) => any]): T[];
    remove(item: T): boolean;
    add(item: T): void;
    addRange(items: T[]): void;
    removeRange(items: T[]): void;
    groupBy(key: any): { key: any, values: T[] }[];
    sum(key?: any): T;
    distinct<U>(callbackfn: (value: T, index: number, array: T[]) => U): U[]
  }
}

Array.prototype.sum = function (key?: any) {
  return this.reduce(function (rv, x) {
    return rv + (key ? x[key] : x);
  }, 0);
};

Array.prototype.distinct = function <T, U>(callbackfn: (value: T, index: number, array: T[]) => U): any {
  return [...new Set(this.map(callbackfn))];
};


Array.prototype.groupBy = function (key: any) {
  return this.reduce(function (rv, x) {
    let v = key instanceof Function ? key(x) : x[key];
    let el = rv.find((r: any) => r && r.key === v);
    if (el) {
      el.values.push(x);
    }
    else {
      rv.push({ key: v, values: [x] });
    }
    return rv;
  }, []);
};

Array.prototype.orderBy = function (propertyExpression: (item: any) => any) {
  let result = [] as any[];
  var compareFunction = (item1: any, item2: any): number => {
    if (propertyExpression(item1) > propertyExpression(item2)) return 1;
    if (propertyExpression(item2) > propertyExpression(item1)) return -1;
    return 0;
  }
  for (var i = 0; i < (<Array<any>>this).length; i++) {
    return (<Array<any>>this).sort(compareFunction);

  }
  return result;
}

Array.prototype.orderByDescending = function (propertyExpression: (item: any) => any) {
  let result = [] as any[];
  var compareFunction = (item1: any, item2: any): number => {
    if (propertyExpression(item1) > propertyExpression(item2)) return -1;
    if (propertyExpression(item2) > propertyExpression(item1)) return 1;
    return 0;
  }
  for (var i = 0; i < (<Array<any>>this).length; i++) {
    return (<Array<any>>this).sort(compareFunction);
  }
  return result;
}

Array.prototype.orderByMany = function (propertyExpressions: [(item: any) => any]) {
  let result = [] as any[];
  var compareFunction = (item1: any, item2: any): number => {
    for (var i = 0; i < propertyExpressions.length; i++) {
      let propertyExpression = propertyExpressions[i];
      if (propertyExpression(item1) > propertyExpression(item2)) return 1;
      if (propertyExpression(item2) > propertyExpression(item1)) return -1;
    }
    return 0;
  }
  for (var i = 0; i < (<Array<any>>this).length; i++) {
    return (<Array<any>>this).sort(compareFunction);
  }
  return result;
}

Array.prototype.orderByManyDescending = function (propertyExpressions: [(item: any) => any]) {
  let result = [] as any[];
  var compareFunction = (item1: any, item2: any): number => {
    for (var i = 0; i < propertyExpressions.length; i++) {
      let propertyExpression = propertyExpressions[i];
      if (propertyExpression(item1) > propertyExpression(item2)) return -1;
      if (propertyExpression(item2) > propertyExpression(item1)) return 1;
    }
    return 0;
  }
  for (var i = 0; i < (<Array<any>>this).length; i++) {
    return (<Array<any>>this).sort(compareFunction);
  }
  return result;
}

Array.prototype.addRange = function (items: any[]): void {
  for (var i = 0; i < items.length; i++) {
    (<Array<any>>this).push(items[i]);
  }
}

Array.prototype.add = function (item: any): void {
  (<Array<any>>this).push(item);
}

Array.prototype.remove = function (item: any): boolean {
  let index = (<Array<any>>this).indexOf(item);
  if (index >= 0) {
    (<Array<any>>this).splice(index, 1);
    return true;
  }
  return false;
}

Array.prototype.removeRange = function (items: any[]): void {
  for (var i = 0; i < items.length; i++) {
    (<Array<any>>this).remove(items[i]);
  }
}

Array.prototype.where = function (predicate: (item: any) => boolean) {
  let result = [];
  for (var i = 0; i < (<Array<any>>this).length; i++) {
    let item = (<Array<any>>this)[i];
    if (predicate(item)) {
      result.push(item);
    }
  }
  return result;
}

Array.prototype.firstOrDefault = function (predicate: (item: any) => boolean) {
  for (var i = 0; i < (<Array<any>>this).length; i++) {
    let item = (<Array<any>>this)[i];
    if (predicate(item)) {
      return item;
    }
  }
  return null;
}

export { }
