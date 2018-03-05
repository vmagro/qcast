//
//  TrackListTable.swift
//  QCastDev
//
//  Created by Vinnie Magro on 11/26/17.
//
//  TrackListTable is a common implementation of a table of tracks that optionally have add buttons to add to the current party's queue -
//   the add buttons also update with changes to the queue

import Foundation
import UIKit
import Kingfisher
import Dwifft
import Material
import Libqcast

class TrackListTable: TableView {
  var tracks: [LibqcastTrack] = [] {
    didSet {
      var ids = [String]()
      for track in tracks {
        ids.append(track.id_())
      }
      DispatchQueue.main.async {
        self.differ.rows = ids
      }
    }
  }
  
  var differ: SingleSectionTableViewDiffCalculator<String>!
  var queue: MobileQueueProtocol!
  
  init() {
    super.init(frame: .zero, style: .plain)
    differ = SingleSectionTableViewDiffCalculator<String>(tableView: self, initialRows: [], sectionIndex: 0)
    differ.deletionAnimation = .automatic
    differ.insertionAnimation = .automatic
    
    queue = MobileWrapQueue(LibqcastRethinkPartyService().currentParty().queue())
    
    self.dataSource = self
    self.delegate = self
    self.showsHorizontalScrollIndicator = false
  }
  
  required init?(coder aDecoder: NSCoder) {
    super.init(coder: aDecoder)
  }
}

extension TrackListTable: UITableViewDataSource, UITableViewDelegate {
  override var numberOfSections: Int {
    return 1
  }
  
  func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
    return differ.rows.count
  }
  
  func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
    let cell = tableView.dequeueReusableCell(withIdentifier: "TableViewCell", for: indexPath) as! TableViewCell
    
    let trackID = differ.rows[indexPath.row]
    var track: LibqcastTrack!
    for t in self.tracks {
      if t.id_() == trackID {
        track = t
        break
      }
    }
    if track == nil {
      return cell
    }
    
    cell.textLabel?.text = track.title()
    cell.imageView?.kf.setImage(with: URL(string: track.album().imageURL()), placeholder: Icon.audio?.resize(toHeight: 64)?.resize(toWidth: 64))
    cell.detailTextLabel?.text = "\(track.artistDisplay() ?? "") - \(track.album().name() ?? "")"
    
    let addButton = IconButton(image: Icon.add, tintColor: Color.green.base)
    addButton.addTarget(self, action: #selector(self.addTrack), for: .touchUpInside)
    // TODO: this might be an abuse of tag but it works for now
    addButton.tag = indexPath.row
    addButton.frame = CGRect(x: 0, y: 0, width: 44, height: 44)
    
    // if track is already in the queue change it to a check and make it disabled
    if queue.track(inQueue: track) {
      addButton.image = Icon.check
      addButton.isEnabled = false
    }
    
    cell.accessoryView = addButton
    cell.dividerColor = Color.grey.lighten2
    
    return cell
  }
  
  func tableView(_ tableView: UITableView, heightForRowAt indexPath: IndexPath) -> CGFloat {
    return 64
  }
}

extension TrackListTable {
  @objc
  func addTrack(sender: Material.Button) {
    do {
      let track = tracks[sender.tag]
      try queue.add(track)
      
      // TODO: probably should do this only when we know it worked, but disable the button once we added it to the queue
      sender.image = Icon.check
      sender.isEnabled = false
    } catch {
      // TODO: do something smart here
    }
  }
}
