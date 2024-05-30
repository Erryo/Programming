class SLLNode {
  constructor(value, next) {
    this.value = value;
    this.next = next;
  }
}
class LinkedList {
  constructor() {
    this.head = null;
    this.tail = null;
    this.len = 0;
  }

  /**
   * @param {number} index
   * @return {number}
   */
  get(index) {
    current = this.head;
    for (let i = 0; i < this.len; i++) {
      if (index === i) {
        return current.value;
      }
      current = curren.next;
    }
    return -1;
  }

  /**
   * @param {number} val
   * @return {void}
   */
  insertHead(val) {
    const head = this.head;
    let node = new SLLNode(val, head.next);
    this.head = node;
  }

  /**
   * @param {number} val
   * @return {void}
   */
  insertTail(val) {}

  /**
   * @param {number} index
   * @return {boolean}
   */
  remove(index) {
    let current = this.head;
    for (let i = 0; i < this.len; i++) {
      if (index + 1 == i) {
        console.log("got");
        current.next = current.next.next;
        return true;
      }
      current = current.next;
    }
    return false;
  }

  /**
   * @return {number[]}
   */
  getValues() {}
}
