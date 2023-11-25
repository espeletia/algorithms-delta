package diary

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type Node struct {
	Text     string
	Time     time.Time
	Next     *Node
	Previous *Node
}

type DiaryEntry struct {
	Text string
	Time time.Time
}

type Diary struct {
	Entries []DiaryEntry
	Head    *Node
	First   *Node
	Last    *Node
}

func NewDiary() *Diary {
	return &Diary{
		Head:  nil,
		First: nil,
		Last:  nil,
	}
}

func (d *Diary) Add(text string) {
	node := &Node{
		Text: text,
		Time: time.Now(),
	}

	if d.Head == nil {
		d.Head = node
		d.First = node
		d.Last = node
		d.Entries = append(d.Entries, DiaryEntry{Text: text, Time: node.Time})
		return
	}

	d.Head.Next = node
	node.Previous = d.Head
	d.Head = node
	d.Entries = append(d.Entries, DiaryEntry{Text: text, Time: node.Time})
}

func (d *Diary) Delete() {
	if d.Head == nil {
		return
	}

	if d.Head == d.First && d.Head == d.Last {
		d.Head = nil
		d.First = nil
		d.Last = nil
		d.Entries = nil
		return
	}

	if d.Head == d.First {
		d.Head = d.Head.Next
		d.Head.Previous = nil
		d.First = d.Head
		d.Entries = d.Entries[1:]
		return
	}

	if d.Head == d.Last {
		d.Head = d.Head.Previous
		d.Head.Next = nil
		d.Last = d.Head
		d.Entries = d.Entries[:len(d.Entries)-1]
		return
	}

	d.Head.Previous.Next = d.Head.Next
	d.Head.Next.Previous = d.Head.Previous
	d.Head = d.Head.Next
	d.Entries = d.Entries[1:]
}

func (d *Diary) Next() {
	if d.Head == nil {
		log.Println("Diary empty")
		return
	}
	if d.Head.Next == nil {
		log.Println("There are no more entries")
		return
	}
	d.Head = d.Head.Next
}

func (d *Diary) Previous() {
	if d.Head == nil {
		log.Println("Diary empty")
		return
	}
	if d.Head.Previous == nil {
		log.Println("This is the last entry")
		return
	}

	// Add a log statement to check the state before updating d.Head
	log.Printf("Before Previous - Head: %+v, Previous: %+v\n", d.Head, d.Head.Previous)

	d.Head = d.Head.Previous

	// Add a log statement to check the state after updating d.Head
	log.Printf("After Previous - Head: %+v, Previous: %+v\n", d.Head, d.Head.Previous)
}

func (d *Diary) ToJSON(filename string) error {
	// Convert the Diary entries to JSON
	data, err := json.Marshal(d.Entries)
	if err != nil {
		return err
	}

	// Write the JSON data to a file
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (d *Diary) FromJSON(filename string) error {
	// Read the JSON data from the file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into the Diary entries
	err = json.Unmarshal(data, &d.Entries)
	if err != nil {
		return err
	}

	// Reconstruct the linked list based on the unmarshaled entries
	d.reconstructLinkedList()

	return nil
}

func (d *Diary) DisplayCurrentEntry() error {
	if d.Head == nil {
		log.Println("Diary empty")
		return nil
	}
	log.Printf("Text Entry: %s\n Date: %s\n", d.Head.Text, d.Head.Time.Format("2006-01-02 15:04:05"))
	return nil
}

func (d *Diary) reconstructLinkedList() {
	// Reconstruct the linked list based on the Diary entries
	if len(d.Entries) == 0 {
		d.Head = nil
		d.First = nil
		d.Last = nil
		return
	}

	// Initialize the linked list with the first entry
	d.Head = &Node{
		Text: d.Entries[0].Text,
		Time: d.Entries[0].Time,
	}
	d.First = d.Head

	// Iterate over the remaining entries to reconstruct the linked list
	for i := 1; i < len(d.Entries); i++ {
		node := &Node{
			Text: d.Entries[i].Text,
			Time: d.Entries[i].Time,
		}
		d.Head.Next = node
		node.Previous = d.Head
		d.Head = node
	}

	// Set the Last pointer to the last node in the reconstructed linked list
	d.Last = d.Head
}
