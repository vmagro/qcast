//
//  SpotifyPlayer.swift
//  QCastDev
//
//  Created by Vinnie Magro on 12/9/17.
//

import Foundation
import MediaPlayer
import Libqcast

// allow throwing strings as errors
extension String: Error {}

protocol PlayerUIDelegate {
  func changePlayPause()
  func changeTrack()
}

class SpotifyPlayer: NSObject, LibqcastNativePlayerInterfaceProtocol, SPTAudioStreamingDelegate, SPTAudioStreamingPlaybackDelegate {
  
  static let sharedInstance = SpotifyPlayer()
  
  var isLoggedIn = false
  fileprivate var player: SPTAudioStreamingController!
  fileprivate var watcher: LibqcastNativePlayerWatcherProtocol?
  fileprivate var currentTrack: LibqcastTrack?
  
  var uiDelegate: PlayerUIDelegate?
  
  fileprivate override init() {
    super.init()
    player = SPTAudioStreamingController.sharedInstance()!
    player.delegate = self
    player.playbackDelegate = self
  }
  
  func login(accessToken: String) throws {
    try player.start(withClientId: LibqcastSpotifyClientID)
    player.login(withAccessToken: accessToken)
  }
  
  func audioStreamingDidLogin(_ audioStreaming: SPTAudioStreamingController!) {
    LibqcastRegisterNativePlayer(self)
    isLoggedIn = true
  }
  
  func audioStreaming(_ audioStreaming: SPTAudioStreamingController!, didReceiveError error: Error!) {
    self.watcher?.onError(error)
  }
  
  func audioStreaming(_ audioStreaming: SPTAudioStreamingController!, didStartPlayingTrack trackUri: String!) {
    uiDelegate?.changeTrack()
    uiDelegate?.changePlayPause()
    // TODO: only do this stuff on the first track to be started
    UIApplication.shared.beginReceivingRemoteControlEvents()
    MPRemoteCommandCenter.shared().skipBackwardCommand.isEnabled = false
    MPRemoteCommandCenter.shared().previousTrackCommand.isEnabled = false
    MPRemoteCommandCenter.shared().seekBackwardCommand.isEnabled = false
    do {
      try AVAudioSession.sharedInstance().setActive(true)
      try AVAudioSession.sharedInstance().setCategory(AVAudioSessionCategoryPlayback, with: [])
    } catch {
      // TODO
    }
    
    // update the notification
    if currentTrack == nil {
      // TODO
      return
    }
    let trackInfo : [String:Any] = [
      MPMediaItemPropertyTitle:  currentTrack!.title(),
      MPMediaItemPropertyArtist: currentTrack!.artistDisplay(),
//      MPMediaItemPropertyArtwork: albumArt,
//      MPNowPlayingInfoPropertyElapsedPlaybackTime: player.
//      MPMediaItemPropertyPlaybackDuration: self.player!.currentTrackDuration,
      MPNowPlayingInfoPropertyPlaybackRate: 1.0,
      MPMediaItemPropertyPlaybackDuration: currentTrack?.duration(),
    ]
    MPNowPlayingInfoCenter.default().nowPlayingInfo = trackInfo
  }
  
  func audioStreaming(_ audioStreaming: SPTAudioStreamingController!, didChangePosition position: TimeInterval) {
    MPNowPlayingInfoCenter.default().nowPlayingInfo![MPNowPlayingInfoPropertyElapsedPlaybackTime] = position
  }
  
  func audioStreaming(_ audioStreaming: SPTAudioStreamingController!, didStopPlayingTrack trackUri: String!) {
    watcher?.onTrackEnd()
    uiDelegate?.changeTrack()
    uiDelegate?.changePlayPause()
  }
  
  func audioStreaming(_ audioStreaming: SPTAudioStreamingController!, didChangePlaybackStatus isPlaying: Bool) {
    uiDelegate?.changePlayPause()
  }
  
  var isPlaying: Bool {
    get {
      if player == nil || player.playbackState == nil {
        return false
      }
      return player.playbackState.isPlaying
    }
  }
  
  func play() throws {
    if player == nil {
      throw "player is nil"
    }
    player.setIsPlaying(true) { (err) in
      if err != nil {
        self.watcher?.onError(err)
      }
    }
  }
  
  func pause() throws {
    if player == nil {
      throw "player is nil"
    }
    player.setIsPlaying(false) { (err) in
      if err != nil {
        self.watcher?.onError(err)
      }
    }
  }
  
  func setTrack(_ track: LibqcastTrack!) throws {
    if player == nil {
      throw "player is nil"
    }
    currentTrack = track
    player.playSpotifyURI(track.playableURL(), startingWith: 0, startingWithPosition: 0) { (err) in
      if err != nil {
        self.watcher?.onError(err)
      }
    }
  }
  
  func notify(_ watcher: LibqcastNativePlayerWatcherProtocol!) {
    self.watcher = watcher
  }
  
  func togglePlayPause() {
    let desired = !isPlaying
    do {
      if desired {
        try self.play()
      } else {
        try self.pause()
      }
    } catch {
      
    }
  }
  
  func skip() {
    // make sure that repeat is off before we skip
    player.setRepeat(.off) { (err) in
      if err != nil {
        self.watcher?.onError(err)
      }
    }
    // TODO: this should probably be somewhere else
    do {
      try LibqcastRethinkPartyService().currentParty().queue().removeCurrent()
    } catch {
      
    }
    player.skipNext { (err) in
      if err != nil {
        self.watcher?.onError(err)
      }
    }
  }
  
}
